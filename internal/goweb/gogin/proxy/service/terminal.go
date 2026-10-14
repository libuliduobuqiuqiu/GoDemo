package service

import (
	"godemo/internal/goweb/gogin/proxy/public"
	"io"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

type TerminalService struct{}

func NewTerminalService() *TerminalService {
	return &TerminalService{}
}

func (t *TerminalService) BuildTerminal(params public.ProxyParams, ws *websocket.Conn) (err error) {
	// 打开Ssh Client
	client, err := newSshClient(params.Address, params.Username, params.Password)
	if err != nil {
		public.HandleErrMessage(ws, err)
		return
	}

	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		public.HandleErrMessage(ws, err)
		return
	}

	defer session.Close()

	if err = session.RequestPty("xterm", 768, 1024, ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}); err != nil {
		public.HandleErrMessage(ws, err)
		return
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		public.HandleErrMessage(ws, err)
		return
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		public.HandleErrMessage(ws, err)
		return
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		public.HandleErrMessage(ws, err)
		return
	}

	go io.Copy(wsWriter(ws), stdout)
	go io.Copy(wsWriter(ws), stderr)
	go io.Copy(stdin, wsReader(ws))

	if err = session.Shell(); err != nil {
		public.HandleErrMessage(ws, err)
		return
	}

	if err = session.Wait(); err != nil {
		logrus.Warnf("session ended with error: %v", err)
	}

	return nil
}

func newSshClient(address, username, password string) (client *ssh.Client, err error) {
	hostKeyCallback, err := HostKeyCallback()
	if err != nil {
		return nil, err
	}
	clientConfig := ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
			setKeyboardInteractive(password),
		},
		HostKeyCallback: hostKeyCallback,
		Timeout:         10 * time.Second,
	}

	client, err = ssh.Dial("tcp", address, &clientConfig)
	if err != nil {
		return
	}

	return
}

func wsReader(ws *websocket.Conn) io.Reader {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				return
			}
			pw.Write(msg)
		}
	}()

	return pr
}

func wsWriter(ws *websocket.Conn) io.Writer {
	pr, pw := io.Pipe()
	go func() {
		defer pr.Close()
		buf := make([]byte, 1024)
		for {
			n, err := pr.Read(buf)
			if err != nil {
				return
			}

			ws.WriteMessage(websocket.TextMessage, buf[:n])
		}
	}()
	return pw
}

// 设置键盘交互
func setKeyboardInteractive(password string) ssh.AuthMethod {
	keyboardInteractiveChallenge := func(
		user,
		instruction string,
		questions []string,
		echos []bool,
	) (answers []string, err error) {
		if len(questions) == 0 {
			return []string{}, nil
		}
		return []string{password}, nil
	}
	return ssh.KeyboardInteractive(keyboardInteractiveChallenge)
}
