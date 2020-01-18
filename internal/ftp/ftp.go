package ftp

import (
	"github.com/pkg/errors"
	"github.com/pkg/sftp"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

// Its a implementation of SFTP protocol for an existing SSH server
var log = logrus.New()


type SFTP struct {
	ftp *sftp.Client
}

func Connection(client *ssh.Client, connStr string) (*SFTP, error) {
	log.Infof("connecting to: %s", connStr)
	c, err := sftp.NewClient(client)
	if err != nil {
		return nil, errors.Wrap(err, "sftp client")
	}

	return &SFTP{
		ftp: c,
	}, nil
}

func (s *SFTP) Create() {
	s.getPWD()

}
func (s *SFTP) getPWD() {
	pwd, err := s.ftp.Getwd()
	if err != nil {
		log.Error(errors.Cause(err))
	}
	log.Infof("pwd: %s", pwd)
}
