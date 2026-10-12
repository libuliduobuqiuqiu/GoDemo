package pkg

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	MysqlType     = "mysql"
	ConfigPathEnv = "GO_DEMO_CONFIG"
)

type GlobalConfig struct {
	SSHConfig      BaseConfig   `json:"ssh"`
	F5Config       BaseConfig   `json:"f5"`
	FTPConfig      BaseConfig   `json:"ftp"`
	MysqlConfig    MysqlConfig  `json:"mysql"`
	CompanyMysql   MysqlConfig  `json:"company"`
	Company57Mysql MysqlConfig  `json:"company57"`
	Cookie         string       `json:"cookie"`
	Devices        []BaseConfig `json:"devices"`
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

// GetGlobalConfig reads an explicit file path or GO_DEMO_CONFIG when the path is empty.
func GetGlobalConfig(configPath string) (GlobalConfig, error) {
	if configPath == "" {
		configPath = os.Getenv(ConfigPathEnv)
	}
	if configPath == "" {
		return GlobalConfig{}, fmt.Errorf("config path is empty: pass a path or set %s", ConfigPathEnv)
	}

	var globalConfig GlobalConfig
	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		return GlobalConfig{}, fmt.Errorf("read config %q: %w", configPath, err)
	}

	if err := json.Unmarshal(fileContent, &globalConfig); err != nil {
		return GlobalConfig{}, fmt.Errorf("parse config %q: %w", configPath, err)
	}
	return globalConfig, nil
}

func GenMysqlDSN(configPath string) (string, error) {
	config, err := GetGlobalConfig(configPath)
	if err != nil {
		return "", err
	}
	mysqlConfig := config.MysqlConfig
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Prefix), nil
}
