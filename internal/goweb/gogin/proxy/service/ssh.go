package service

import (
	"bytes"
	"fmt"
	"godemo/internal/goweb/gogin/proxy/public"
	"time"

	"golang.org/x/crypto/ssh"
)

type SshService struct{}

func NewSshService() *SshService {
	return &SshService{}
}

func (s *SshService) ExecCommand(params public.ProxySshParams) (cmdRes map[string]string, err error) {
	client, err := sshConnect(params.Address, params.Username, params.Password, params.Port, true)
	if err != nil {
		return
	}

	defer client.Close()
	cmdRes = make(map[string]string)
	for _, cmd := range params.Commands {
		session, err := client.NewSession()
		if err != nil {
			return nil, err
		}
		defer session.Close()

		var stdoutBuf, stderrBuf bytes.Buffer
		session.Stdout = &stdoutBuf
		session.Stderr = &stderrBuf
		if err = session.Run(cmd); err != nil {
			err = fmt.Errorf("Run command:%s:%s, error: %w", cmd, stderrBuf.String(), err)
			return nil, err
		}
		cmdRes[cmd] = stdoutBuf.String()
	}

	return
}

func sshConnect(address, username, password string, port int, isKeyboardInteractive bool) (client *ssh.Client, err error) {
	hostKeyCallback, err := HostKeyCallback()
	if err != nil {
		return nil, err
	}
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: hostKeyCallback,
		Timeout:         10 * time.Second,
	}

	if isKeyboardInteractive {
		config.Auth = []ssh.AuthMethod{setKeyboardInteractive(password)}
	}

	if port != 0 {
		address = fmt.Sprintf("%s:%d", address, port)
	} else {
		address = fmt.Sprintf("%s:22", address)
	}

	client, err = ssh.Dial("tcp", address, config)
	if err != nil {
		return
	}

	return
}
