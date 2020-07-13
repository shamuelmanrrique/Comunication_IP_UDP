package functions

import (
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func InitSSH(user string, addr string, idRsa string) (ssh.Session, error) {
	key, err := ioutil.ReadFile(idRsa)
	var session *ssh.Session
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", net.JoinHostPort(addr, "22"), config)
	if err != nil {
		panic(err.Error())
	}

	// Create a session. It is one session per command.
	session, err = client.NewSession()
	if err != nil {
		panic(err.Error())
	}

	return *session, err

}
