package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/client"
	transportssh "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/crypto/ssh"
)

func run() error {

	publicKeys, err := transportssh.NewPublicKeysFromFile("git", "/Users/twp/.ssh/id_rsa", "Rib tickler")
	if err != nil {
		return err
	}

	sshClientConfig, err := publicKeys.ClientConfig()
	if err != nil {
		return err
	}
	sshClientConfig.HostKeyCallback = func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		log.Printf("hostname=%s remote=%s key=%s", hostname, remote, key)
		return nil
	}
	client.InstallProtocol("ssh", transportssh.NewClient(sshClientConfig))

	tempDir, err := os.MkdirTemp("", "test-ssh")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	cloneOptions := &git.CloneOptions{
		Auth:  publicKeys,
		URL:   "git@github.com:twpayne/dotfiles.git",
		Depth: 1,
	}
	_, err = git.PlainClone(tempDir, false, cloneOptions)
	return err
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
