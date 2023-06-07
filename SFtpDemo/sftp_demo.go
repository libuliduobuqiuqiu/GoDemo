package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

type ServerConfig struct {
	Host     string `json:"host"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Port     int    `json:"port"`
}

func sftp_connect(serverInfo *ServerConfig) {

	keyboardInteractiveChallenge := func(
		user,
		instruction string,
		questions []string,
		echos []bool,
	) (answers []string, err error) {
		if len(questions) == 0 {
			return []string{}, nil
		}
		return []string{serverInfo.PassWord}, nil
	}

	config := &ssh.ClientConfig{
		User: "admin",
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(keyboardInteractiveChallenge),
			ssh.Password(serverInfo.PassWord),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%v", serverInfo.Host, serverInfo.Port), config)
	if err != nil {
		log.Fatalf("Failed to connect to remote host: %s", err)
	}
	defer conn.Close()

	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("Failed to initialize SFTP client: %s", err)
	}
	defer client.Close()

	remoteDir := "/config"
	files, err := client.ReadDir(remoteDir)
	if err != nil {
		log.Fatalf("Failed to traverse remote directory: %s", err)
	}

	for _, file := range files {

		if file.IsDir() {
			fmt.Println("Dir: ", file.Name())
		}
		fmt.Println(file.Name())

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
		sftp_connect(&val)
	}

}
