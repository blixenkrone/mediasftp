package server

import (
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

type SSHServer struct {
	session *ssh.Session
	client  *ssh.Client
}

var hostKey ssh.PublicKey

func NewSSH(host, port, user, password string) (*SSHServer, error) {
	connStr := fmt.Sprintf("%v:%v", host, port)

	config := &ssh.ClientConfig{
		Config: ssh.Config{},
		User:   user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
		Timeout:         1000,
	}

	sshClient, err := ssh.Dial("tcp", connStr, config)
	if err != nil {
		return nil, errors.Wrap(err, "ssh dial tcp")
	}

	session, err := sshClient.NewSession()
	if err != nil {
		return nil, errors.Wrap(err, "ssh client session")
	}

	return &SSHServer{
		session: session,
		client:  sshClient,
	}, nil
}

func (s *SSHServer) Close() error {
	return s.client.Close()
}
