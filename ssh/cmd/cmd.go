package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

type LoginInfo struct {
	User     string
	Ip       string
	Port     int
	Password string
}

type Client struct {
	*ssh.Client
}

// 在一次session中可连续执行Run方法
type Stream struct {
	in              io.WriteCloser //session.StdinPipe()
	out             *bytes.Buffer  //记录session.Stdout和session.Stderr
	ch              chan string    //保存Run操作后的返回值
	readUntilExpect string         //Run操作执行后读取返回值直到readUntilExpect
	session         *ssh.Session
	logger          *logrus.Logger //记录日志
	mu              *sync.Mutex    //读写锁，确保一次完整的Run操作
}

func NewClient(li *LoginInfo) (*Client, error) {
	config := &ssh.ClientConfig{
		Timeout: time.Second * 5,
		User:    li.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(li.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", li.Ip+":"+strconv.Itoa(li.Port), config)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

// 单次
func (c *Client) Run(cmd string) string {
	session, err := c.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	res, err := session.Output(cmd)
	if err != nil {
		log.Panic(err)
	}
	return string(res)
}

// 新建一个stream
func (c *Client) NewStream() (*Stream, error) {
	// 确定readUntilExpect
	var flag string
	if c.Client.User() == "root" {
		flag = "]#"
	} else {
		flag = "]$"
	}
	// 定义日志
	var log = logrus.New()
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	log.Out = file
	// 创建session
	session, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	if err = session.RequestPty("xterm", 80, 40, modes); err != nil {
		fmt.Printf("get pty error:%v\n", err)
		return nil, err
	}
	stream, err := session.StdinPipe()
	if err != nil {
		log.Panicf("get stdin pipe error%v\n", err)
		return nil, err
	}
	var outbuf = bytes.NewBuffer(make([]byte, 0))
	session.Stdout = outbuf
	session.Stderr = outbuf
	err = session.Shell()
	if err != nil {
		log.Panicf("shell session error%v", err)
	}
	// 过滤返回的登录信息(Last login: Fri Dec  2 08:09:12 2022 from 192.168.101.105),返回stream
	timeout := time.After(time.Second * 10)
	for {
		select {
		case <-timeout:
			log.Panic("stream create timeout")
		default:
			time.Sleep(time.Microsecond * 200)
			if strings.Contains(outbuf.String(), flag) {
				tmp := make([]byte, len(outbuf.String()))
				outbuf.Read(tmp)
				return &Stream{in: stream,
					out:             outbuf,
					ch:              make(chan string, 1),
					session:         session,
					readUntilExpect: flag,
					logger:          log,
					mu:              new(sync.Mutex)}, nil
			}
		}
	}
}

// 更新readUntilExpect
func (s *Stream) UpdateReadUntilExpect(expect string) {
	s.readUntilExpect = expect
}

// 读取返回值，直到readUntilExpect
func (s *Stream) readUntil() error {
	ch := make(chan struct{}, 1)
	timeout := time.After(time.Second * 10)
	for {
		select {
		case <-timeout:
			return errors.New("timeout Waiting for Return")
		case <-ch:
			return nil
		default:
			time.Sleep(time.Microsecond * 200)
			out := s.out.String()
			if strings.Contains(out, s.readUntilExpect) {
				tmp := make([]byte, len(out))
				s.out.Read(tmp)
				s.ch <- string(tmp)
				close(ch)
			}
		}
	}
}

// 每次输入cmd,返回对应的值
func (s *Stream) Run(cmd string) string {
	s.mu.Lock()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		s.logger.Infof("Input:%s", cmd)
		s.in.Write([]byte(fmt.Sprintf("%v\n", cmd)))
		wg.Done()
	}()
	go func() {
		err := s.readUntil()
		if err != nil {
			log.Panic(err)
		}
		wg.Done()
	}()
	wg.Wait()
	out := <-s.ch
	// 返回值过滤掉主机名等信息
	if strings.Contains(out, "]#") || strings.Contains(out, "]$") {
		outSlice := strings.Split(out, "\n")
		outStrip := strings.Join(outSlice[:len(outSlice)-1], "\n")
		out = strings.ReplaceAll(outStrip, "\r", "")
	}
	s.logger.Infof("Output:%s", out)
	s.mu.Unlock()
	return out
}

// 关闭stream
func (s *Stream) Close() {
	s.in.Close()
	s.session.Close()
}
