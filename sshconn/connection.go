package sshconn

import (
	"log"
	"regexp"

	"code.google.com/p/go.crypto/ssh"
)

// Connection contains the necessary configuration for SSH to the tmate proxy
type Connection struct {
	User string
	Host string
}

// NewConnection creates a struct for an SSH connection
func NewConnection(user, host string) *Connection {
	match, err := regexp.MatchString(":", host)
	if err != nil {
		log.Fatalf("unable to match string: %s", err)
	}
	if !match {
		host += ":22"
	}
	return &Connection{User: user, Host: host}
}

// Interactive initiates the interactive SSH session with the terminal
// https://godoc.org/code.google.com/p/go.crypto/ssh#Session.RequestPty
func (sshconn Connection) Interactive() (err error) {
	config := &ssh.ClientConfig{
		User: sshconn.User,
	}
	// Connect to ssh server
	conn, err := ssh.Dial("tcp", sshconn.Host, config)
	if err != nil {
		log.Fatalf("unable to connect: %s", err)
	}
	defer conn.Close()
	// Create a session
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("unable to create session: %s", err)
	}
	defer session.Close()
	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// Request pseudo terminal
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Fatalf("request for pseudo terminal failed: %s", err)
	}
	// Start remote shell
	if err := session.Shell(); err != nil {
		log.Fatalf("failed to start shell: %s", err)
	}
	return
}
