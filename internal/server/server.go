package server

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

type SSHServer struct {
	session *ssh.Session
	client  *ssh.Client
}

var (
	log     = logrus.New()
	hostKey ssh.PublicKey
)

func NewSSH(host, port, user, password string) (*SSHServer, error) {
	connStr := fmt.Sprintf("%v:%v", host, port)
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	sshClient, err := ssh.Dial("tcp", connStr, config)
	if err != nil {
		log.Error(errors.Cause(err))
		return nil, errors.Wrap(err, "ssh dial tcp")
	}

	session, err := sshClient.NewSession()
	if err != nil {
		return nil, errors.Wrap(err, "ssh client session")
	}

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

	return &SSHServer{
		session: session,
		client:  sshClient,
	}, nil
}

type Conn interface {
	Close() error
}

func (s *SSHServer) Close() error {
	var err error
	err = s.client.Close()
	if err != nil {
		return err
	}
	err = s.session.Close()
	if err != nil {
		return err
	}
	return err
}
