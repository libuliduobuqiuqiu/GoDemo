package GoSshDemo

import (
	"bytes"
	"fmt"
	gossh "golang.org/x/crypto/ssh"
	"log"
	"sunrun/ConfigParserDemo"
)

func ExecServerCommand() {

	globalConfig := ConfigParserDemo.GetGlobalConfig()
	config := globalConfig.SSHConfig

	clientConfig := &gossh.ClientConfig{
		User: config.Username,
		Auth: []gossh.AuthMethod{
			gossh.Password(config.Password),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
	}

	client, err := gossh.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port), clientConfig)
	if err != nil {
		log.Fatal("Failed to login: ", err)
	}

	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}

	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b

	if err := session.Run("bash -c -i 'history -r; history'"); err != nil {
		log.Fatal("Failed to run ", err.Error())
	}

	fmt.Println(b.String())

}
