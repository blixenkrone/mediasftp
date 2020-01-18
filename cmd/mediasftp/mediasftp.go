package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"

	"github.com/blixenkrone/mediasftp/internal/server"
)

var (
	host     = flag.String("host", "", "set the hostname for the ftp server")
	port     = flag.String("port", "22", "set the port for the ftp server")
	user     = flag.String("user", "", "set the user for the ftp server")
	password = flag.String("password", "", "set the password for the ftp server")
)

func main() {
	if err := parseArgs(); err != nil {
		panic(err)
	}
	s, err := server.NewSSH(*host, *port, *user, *password)
	if err != nil {
		panic(err)
	}
	defer func() {
		panic(s.Close())
	}()

}

func parseArgs() error {
	flag.Parse()

	if len(os.Args[:0]) < 0 {
		flag.PrintDefaults()
		return errors.New("no flags provided")
	}
	if *user == "" || *password == ""{
		return errors.New("no/wrong user or password provided")
	}
	for i := len(os.Args) - 1; i > 0; i-- {
		fmt.Println(os.Args[i])
	}
	return nil
}
