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

func SFTPConnection(client *ssh.Client) (*SFTP, error) {
	log.Infof("connecting to: %s", client.LocalAddr().String())
	c, err := sftp.NewClient(client)
	if err != nil {
		return nil, errors.Wrap(err, "sftp client")
	}
	return &SFTP{
		ftp: c,
	}, nil
}

func (s *SFTP) Ping() {
	s.getPWD()
	s.ftp.Lock()
	defer s.ftp.Unlock()
	log.Info("PONG!")

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
