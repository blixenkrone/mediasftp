package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {

	user := ""
	pass := ""
	remote := ""
	port := ":22"

	// get host public key
	hostKey := getHostKey(remote)

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
		HostKeyAlgorithms: []string{
			ssh.KeyAlgoRSA,
		},
	}
	// connect
	conn, err := ssh.Dial("tcp", remote+port, config)
	if err != nil {
		log.Fatal(errors.Wrap(err, "ssh dial error:"))
	}
	defer conn.Close()

	// create new SFTP client
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// create destination file
	dstFile, err := client.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// create source file
	srcFile, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)
}

func getHostKey(host string) ssh.PublicKey {
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "authorized_keys"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Println(file.Name())

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		b := scanner.Bytes()
		hostKey, _, _, _, err = ssh.ParseAuthorizedKey(b)
		if err != nil {
			panic(err)
		}
		break
	}



	// for scanner.Scan() {
	// 	fields := strings.Split(scanner.Text(), " ")
	// 	if len(fields) != 3 {
	// 		continue
	// 	}
	// 	if strings.Contains(fields[0], host) {
	// 		var err error
	// 		hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
	// 		if err != nil {
	// 			log.Fatalf("error parsing %q: %v", fields[2], err)
	// 		}
	// 		break
	// 	}
	// }

	if hostKey == nil {
		log.Fatalf("no hostkey found for %s", host)
	}

	spew.Dump(string(hostKey.Marshal()))
	return hostKey
}
