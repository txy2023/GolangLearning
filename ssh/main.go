package main

import (
	"io"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func SendCommand(in io.WriteCloser, cmd string) error {
	if _, err := in.Write([]byte(cmd + "\n")); err != nil {
		return err
	}

	return nil
}

func main() {

	// Setup configuration for SSH client
	config := &ssh.ClientConfig{
		Timeout: time.Second * 5,
		User:    "tian",
		Auth: []ssh.AuthMethod{
			ssh.Password("tian"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the client
	client, err := ssh.Dial("tcp", "192.168.101.109:22", config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Create a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Setup StdinPipe to send commands
	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	// defer stdin.Close()

	// Route session Stdout/Stderr to system Stdout/Stderr
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// Start a shell
	if err := session.Shell(); err != nil {
		log.Fatal(err)
	}

	// Send username
	if _, err := stdin.Write([]byte("pwd\n")); err != nil {
		log.Fatal(err)
	}
	// SendCommand(stdin, "cd /home\n")
	SendCommand(stdin, "pwd")
	SendCommand(stdin, "su")
	SendCommand(stdin, "tian")
	SendCommand(stdin, "whoami")
	SendCommand(stdin, "whoami")
	SendCommand(stdin, "echo 'echo hello world'>test1.sh && chmod 777 test1.sh")
	SendCommand(stdin, "./test1.sh")
	stdin.Close()
	err = session.Wait()
	if err != nil {
		log.Println(err)
	}
}
