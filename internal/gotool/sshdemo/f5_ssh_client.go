package sshdemo

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"sunrun/pkg"
)

func Connect(s pkg.BaseConfig, cmd string) {
	// SSH连接信息
	keyboardInteractiveChallenge := func(
		user,
		instruction string,
		questions []string,
		echos []bool,
	) (answers []string, err error) {
		if len(questions) == 0 {
			return []string{}, nil
		}
		return []string{s.Password}, nil
	}

	config := &ssh.ClientConfig{
		User: s.Username,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(keyboardInteractiveChallenge),
			ssh.Password(s.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 建立SSH连接
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%v", s.Host, s.Port), config)
	if err != nil {
		fmt.Printf("Failed to dial: %s\n", err)
		return
	}
	defer conn.Close()

	// 执行命令
	session, err := conn.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s\n", err)
		return
	}
	defer session.Close()

	// 执行命令
	output, err := session.CombinedOutput(cmd)

	fmt.Println(string(output))
	if err != nil {
		fmt.Printf("Failed to run command: %s\n", err)
		return
	}
}

func ExecF5Command() {
	f5Config := pkg.GetGlobalConfig()
	config := f5Config.F5Config
	cmd := "tmsh list sys softwareddd;\nifconfig lo"
	Connect(config, cmd)

}
