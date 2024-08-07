package pkg

import (
	"encoding/json"
	"log"
	"os"
)

const ConfigPath = "/data/GoDemo/configs/conf.json"

type GlobalConfig struct {
	SSHConfig      BaseConfig  `json:"ssh"`
	F5Config       BaseConfig  `json:"f5"`
	FTPConfig      BaseConfig  `json:"ftp"`
	MysqlConfig    MysqlConfig `json:"mysql"`
	CompanyMysql   MysqlConfig `json:"company"`
	Company57Mysql MysqlConfig `json:"company57"`
}

type MysqlConfig struct {
	BaseConfig
	Prefix string `json:"prefix"`
}

type BaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func GetGlobalConfig() GlobalConfig {
	var globalConfig GlobalConfig
	if _, err := os.Stat(ConfigPath); err != nil {
		log.Fatal(err)
	}

	fileContent, err := os.ReadFile(ConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(fileContent, &globalConfig); err != nil {
		log.Fatal()
	}
	return globalConfig
}
