package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

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

type Stream struct {
	in      io.WriteCloser
	out     *bytes.Buffer
	ch      chan string
	session *ssh.Session
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

	// Connect to the client
	client, err := ssh.Dial("tcp", li.Ip+":"+strconv.Itoa(li.Port), config)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

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

func (c *Client) NewStreamPipe() (*Stream, error) {
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
		log.Printf("get stdin pipe error%v\n", err)
		return nil, err
	}
	var outbuf *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
	session.Stdout = outbuf
	session.Stderr = outbuf

	err = session.Shell()
	if err != nil {
		fmt.Printf("shell session error%v", err)
		return nil, err
	}
	go session.Wait()
	return &Stream{in: stream, out: outbuf, ch: make(chan string, 1), session: session}, nil
}

func (s *Stream) Run(cmd string) string {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		s.in.Write([]byte(fmt.Sprintf("%v\n", cmd)))
		wg.Done()
	}()
	go func() {
		buf := make([]byte, 8192)
		var t int
		terminator := '$'
		for {
			fmt.Println("test")
			time.Sleep(time.Second * 1)
			n, err := s.out.Read(buf)
			if err != nil && err != io.EOF {
				log.Panic(err)
				fmt.Println(string(buf))
			}
			if n > 0 {
				t = bytes.LastIndexByte(buf, byte(terminator))
				if t > 0 {
					s.ch <- string(buf[:t])
					break
				} else {
					s.ch <- string(buf)
					break
				}
			}
		}
		wg.Done()
	}()
	wg.Wait()
	return <-s.ch
}

func (s Stream) Close() {
	s.in.Close()
	s.session.Close()

}
