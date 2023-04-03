package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
)

type ServerConfig struct {
	Host     string `json:"host"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Port     int    `json:"port"`
}

func Connect(s *ServerConfig, cmd string) {
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
		return []string{s.PassWord}, nil
	}

	config := &ssh.ClientConfig{
		User: s.UserName,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(keyboardInteractiveChallenge),
			ssh.Password(s.PassWord),
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

func main() {
	config := make(map[string]ServerConfig)

	confPath := "conf.json"
	if _, err := os.Stat(confPath); err != nil {
		fmt.Printf("Config File Error: %s", err)
		return
	}

	content, err := os.ReadFile(confPath)
	if err != nil {
		fmt.Printf("Error Reading File: %s", err)
		return
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		fmt.Printf("Error Unmarshal: %s", err)
		return
	}

	if val, ok := config["F5"]; ok {
		cmd := "tmsh list sys softwareddd;\nifconfig lo"
		Connect(&val, cmd)
	}

}
