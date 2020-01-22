package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/blixenkrone/mediasftp/internal/ftp"
	"github.com/blixenkrone/mediasftp/internal/ssh"
)

var (
	log      = logrus.New()
	env      = flag.String("env", "development", "set env")
	host     = flag.String("host", "", "set the hostname for the ftp server")
	port     = flag.String("port", "22", "set the port for the ftp server")
	user     = flag.String("user", "", "set the user for the ftp server")
	password = flag.String("password", "", "set the password for the ftp server")
)

func main() {
	if err := parseArgs(*env); err != nil {
		panic(err)
	}
	serverConnStr := fmt.Sprintf("%v:%v", *host, *port)
	log.Info(serverConnStr)
	l, err := net.Listen("tcp", *host+":"+*port)
	if err != nil {
		panic(err)
	}
	defer func() {
		panic(l.Close())
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
			return
		}
		res := make(chan<- string)
		errCh := make(chan<- error)

		// req, err := newRequest(os.Getenv("USER"), os.Getenv("PASSWORD"))

		req := &TCPRequest{}

		if err != nil {
			panic(err)
			return
		}
		go req.handleConnection(conn, res, errCh)
	}

}

type TCPRequest struct {
	Client *ftp.SFTP
}

func newRequest(user, password string) (*TCPRequest, error) {
	host, port := os.Getenv("HOST_IP"), os.Getenv("HOST_PORT")
	c, err := ssh.Client(host, port, user, password)
	if err != nil {
		return nil, err
	}
	sftp, err := ftp.SFTPConnection(c.Client)
	return &TCPRequest{
		Client: sftp,
	}, nil
}

func (sftp *TCPRequest) handleConnection(c net.Conn, res chan<- string, err chan<- error) {
	c.Write([]byte("hello"))
}

func parseArgs(env string) error {
	flag.Parse()
	fName := fmt.Sprintf("%s.env", env)
	if err := godotenv.Load(fName); err != nil {
		return err
	}

	log.Info(env)

	switch env {
	case "local":
		*host = "localhost"
		*port = "2222"
		break

	case "development":

		break
	default:
		return errors.New("environment var not provided")
		break
	}

	if len(os.Args[:1]) < 0 {
		flag.PrintDefaults()
		return errors.New("no flags provided")
	}
	return nil
}
