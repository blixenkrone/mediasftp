package ssh

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

var (
	log     = logrus.New()
	hostKey ssh.PublicKey
)

type SSH struct {
	Client *ssh.Client
	Config *clientCfg
}

type clientCfg struct {
	host, port, user, password string
}

func Client(host, port, user, password string) (*SSH, error) {
	cfg := &clientCfg{
		host:     host,
		port:     port,
		user:     user,
		password: password,
	}
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		// HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	client, err := ssh.Dial("tcp", cfg.GetConnStr(), config)
	if err != nil {
		log.Error(errors.Cause(err))
		return nil, errors.Wrap(err, "ssh dial tcp")
	}

	return &SSH{
		Client: client,
		Config: cfg,
	}, nil
}

func getHostKey(host string) ssh.PublicKey {
	// parse OpenSSH known_hosts file
	// ssh or use ssh-keyscan to get initial key
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}

	if hostKey == nil {
		log.Fatalf("no hostkey found for %s", host)
	}

	return hostKey
}

type Conn interface {
	Close() error
}

func (s *SSH) Close() error {
	return s.Client.Close()
}

func (cfg *clientCfg) GetConnStr() string {
	return cfg.host + ":" + cfg.port
}
