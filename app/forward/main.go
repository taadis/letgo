package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

// get default location of a private key
func privateKeyPath() string {
	return os.Getenv("HOME") + "/.ssh/id_rsa"
}

// Get private key for ssh authentication
func parsePrivateKey(keyPath string) (ssh.Signer, error) {
	filename := keyPath
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("ioutil.ReadFile error:%+v\n", err)
		return nil, err
	}
	return ssh.ParsePrivateKey(buf)
}

// Get ssh client config for our connection
// SSH config will use 2 authentication strategies: by key and password
func makeConfigSSH(user string, password string) (*ssh.ClientConfig, error) {
	keyPath := privateKeyPath()
	key, err := parsePrivateKey(keyPath)
	if err != nil {
		log.Printf("parsePrivateKey error:%+v", err)
		return nil, err
	}

	config := &ssh.ClientConfig{}
	config.User = user
	config.Auth = []ssh.AuthMethod{
		ssh.PublicKeys(key),
		ssh.Password(password),
	}

	return config, nil
}

// Handle local client connections and tunnel data to the remote server
// Will use io.Copy - https://pkg.go.dev/io#Copy
func handleClient(client net.Conn, remote net.Conn) {
	defer client.Close()
	defer remote.Close()

	doneCh := make(chan bool)

	// Start remote -> local data transfer
	go func() {
		_, err := io.Copy(remote, client)
		if err != nil {
			log.Printf("remote -> local io.Copy error:%+v", err)
		}
		doneCh <- true
	}()

	// Start local -> remote data transfer
	go func() {
		_, err := io.Copy(remote, client)
		if err != nil {
			log.Printf("local -> remote io.Copy error:%+v", err)
		}
		doneCh <- true
	}()

	<-doneCh
}

func main() {
	// Connection settings
	sshAddr := "remote-ip:22"
	localAddr := "127.0.0.1:5000"
	remoteAddr := "127.0.0.1:5432"

	// Build SSH client configuration
	cfg, err := makeConfigSSH("user", "password")
	if err != nil {
		log.Fatalf("makeConfigSSH error:%+v\n", err)
		return
	}

	// Establish connection with SSH server
	conn, err := ssh.Dial("tcp", sshAddr, cfg)
	if err != nil {
		log.Fatalf("ssh Dial error:%+v", err)
		return
	}
	defer conn.Close()

	// Establish connection with remote server
	remoteConn, err := conn.Dial("tcp", remoteAddr)
	if err != nil {
		log.Fatalf("conn.Dial error:%+v", err)
		return
	}

	// Start local server to forward traffic to remote connection
	local, err := net.Listen("tcp", localAddr)
	if err != nil {
		log.Fatalf("net.Listen error:%+v", err)
		return
	}
	defer local.Close()

	// Handle incoming connections
	for {
		client, err := local.Accept()
		if err != nil {
			log.Fatalf("local.Accept error:%+v", err)
		}

		handleClient(client, remoteConn)
	}
}
