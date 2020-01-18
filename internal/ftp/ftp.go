package ftp

import (
	"github.com/pkg/errors"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Its a implementation of SFTP protocol for an existing SSH server

type SFTP struct {
	opt    *sftp.ClientOption
	server *sftp.Server
	client *sftp.Client
}

var conn *ssh.Client

func Connection(host string, port int) (*SFTP, error) {
	c, err := sftp.NewClient(conn)
	sftp.Serv
	if err != nil {
		return nil, errors.Wrap(err, "sftp client")
	}

	return &SFTP{
		client: c,
	}, nil

}
