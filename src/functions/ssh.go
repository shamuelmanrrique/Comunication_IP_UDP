package functions

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

func InitSSH(addr string) *ssh.Client {
	// IDRsa := "/home/smmanrrique/.ssh/id_rsa"
	IDRsa := "/home/smmanrrique/.ssh/id_rsa"
	// var user = "a802400"
	var user = "smmanrrique"

	println("aqui en ssh", addr)

	key, err := ioutil.ReadFile(IDRsa)
	if err != nil {
		println("ERRRRRRRROR en ssh", addr)
		panic(err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", net.JoinHostPort(addr, "22"), config)
	if err != nil {
		panic(err.Error())
	}

	return client

}

func ExcecuteSSH(cmd string, conn *ssh.Client) {
	sess, err := conn.NewSession()
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	sessStdOut, err := sess.StdoutPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stdout, sessStdOut)
	sessStderr, err := sess.StderrPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stderr, sessStderr)
	log.Println(cmd)
	err = sess.Run(cmd)
	if err != nil {
		panic(err)
	}
}

func ConnectViaSsh(user, host string, password string) *ssh.Client {
	// func connectViaSsh(user, host string, password string) (*ssh.Client, *ssh.Session) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(SshInteractive),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host, config)
	println(err)
	// session, err := client.NewSession()
	// fmt.Println(err)

	// return client, session
	return client
}

func SshInteractive(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	answers = make([]string, len(questions))
	// The second parameter is unused
	for n, _ := range questions {
		answers[n] = "Hiberus7923"
	}

	return answers, nil
}
