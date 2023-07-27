package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/jlaffaye/ftp"
	"log"
	"os"
	"time"
)

type FtpConfig struct {
	Ftp FtpConn `json:"ftp"`
}

type FtpConn struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func initFtpConn() (err error, ftpConn *ftp.ServerConn) {

	confPath := "conf.json"

	if _, err = os.Stat(confPath); err != nil {
		return
	}

	var Config FtpConfig
	if err = configor.Load(&Config, confPath); err != nil {
		return
	}

	ftpConfig := Config.Ftp
	ftpConn, err = ftp.Dial(fmt.Sprintf("%s:%d", ftpConfig.Host, ftpConfig.Port), ftp.DialWithDisabledEPSV(true))
	if err != nil {
		log.Fatal(err)
		return
	}

	err = ftpConn.Login(ftpConfig.Username, ftpConfig.Password)
	if err != nil {
		return
	}

	return
}

func main() {
	err, ftpConn := initFtpConn()
	if err != nil {
		log.Fatal(err)
		return
	}

	uploadFtpPath := "/home/ftpuser"

	err = ftpConn.ChangeDir(uploadFtpPath)
	if err != nil {
		fmt.Println("mkdir .....")
		err = ftpConn.MakeDir("ftp")
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println("changedir....")
		err = ftpConn.ChangeDir("/home/ftpuser")
		if err != nil {
			log.Println(err.Error())
		}
	}

	// 切换目录
	fmt.Println("list.....")
	//err = ftpConn.ChangeDirToParent()
	//if err != nil {
	//	log.Println(err)
	//}

	// 查看当前目录文件
	files, err := ftpConn.List(".")
	if err != nil {
		log.Println(err.Error())
	}
	for _, file := range files {
		fmt.Println(file.Name)
	}

	//fmt.Println("Upload")
	// 上传文件
	file, err := os.Open("CHANGELOG.md")
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10)
	defer file.Close()
	err = ftpConn.Stor(fmt.Sprintf("%s//CHANGELOG.md", uploadFtpPath), file)
	if err != nil {
		log.Fatal(err)
	}

	err = ftpConn.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
