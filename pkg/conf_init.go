package pkg

import (
	"encoding/json"
	"fmt"
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

func GetGlobalConfig(configPath string) GlobalConfig {
	if configPath == "" {
		configPath = ConfigPath
	}

	var globalConfig GlobalConfig
	if _, err := os.Stat(configPath); err != nil {
		log.Fatal(err)
	}

	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(fileContent, &globalConfig); err != nil {
		log.Fatal()
	}
	return globalConfig
}

func GenMysqlDSN(configPath string) string {
	config := GetGlobalConfig(configPath)
	mysqlConfig := config.MysqlConfig
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Prefix)
}
