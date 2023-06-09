package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"log"
)

func main() {
	ftpConn, err := ftp.Dial("126.0.0.1:21")

	if err != nil {
		log.Fatal(err.Error())
	}

	err = ftpConn.Login("ftp", "123456")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("logout.....")

		errFtp := ftpConn.Logout()
		if errFtp != nil {
			log.Fatal(errFtp.Error())
		}
		fmt.Println("Quit......")
		errFtp = ftpConn.Quit()
		if errFtp != nil {
			log.Fatal(errFtp.Error())
		}
	}()

	err = ftpConn.ChangeDir("ftp")
	if err != nil {
		fmt.Println("mkdir .....")
		err = ftpConn.MakeDir("ftp")
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println("changedir....")
		err = ftpConn.ChangeDir("F://ftp")
		if err != nil {
			log.Println(err.Error())
		}
	}

	fmt.Println("list.....")
	err = ftpConn.ChangeDirToParent()
	if err != nil {
		log.Println(err)
	}

	files, err := ftpConn.List(".")
	if err != nil {
		log.Println(err.Error())
	}
	for _, file := range files {
		fmt.Println(file.Name)
	}

}
