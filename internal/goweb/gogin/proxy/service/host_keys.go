package service

import (
	"errors"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func HostKeyCallback() (ssh.HostKeyCallback, error) {
	path := os.Getenv("GO_DEMO_SSH_KNOWN_HOSTS")
	if path == "" {
		return nil, errors.New("GO_DEMO_SSH_KNOWN_HOSTS is required for SSH routes")
	}
	return knownhosts.New(path)
}
