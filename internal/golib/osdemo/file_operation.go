package osdemo

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/pkg/errors"
)

var filePath = "/data/GoDemo/internal/golib/osdemo/demo.json"
var newFilePath = "/data/GoDemo/internal/golib/osdemo/demo2.json"
var basePath = "/data/GoDemo/internal/golib/osdemo/"

func readFileData(f *os.File) (data string, err error) {

	buffer := make([]byte, 512)
	count, err := f.Read(buffer)
	if err != nil {
		return
	}
	fmt.Println(count)
	data = string(buffer)
	fmt.Println(data)
	return
}

func SimpleOpenFile() error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	readFileData(file)
	return nil
}

func AdvancedOpenFile() error {
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	readFileData(file)
	return nil
}

func SimpleReadFile() error {

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

func AdvancedReadFile() (data []byte, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}

	defer f.Close()

	data = make([]byte, 0, 512)
	for {
		if len(data) >= cap(data) {
			d := append(data[:cap(data)], 0)
			data = d[:len(data)]
		}
		n, err := f.Read(data[len(data):cap(data)])
		data = data[:len(data)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return data, err
		}
	}
}

func IOReadFile() error {
	f, err := os.Open(newFilePath)
	if err != nil {
		return err
	}

	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "读取文件数据失败")
	}

	fmt.Println(string(data))
	return nil
}

func SimpleWriteFile() error {
	file, err := os.OpenFile(newFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	fmt.Println("文件打开成功： ", file.Name())
	for i := 0; i < 5; i++ {
		_, err := file.WriteString("Hello,world\n")
		if err != nil {
			return errors.Wrap(err, "文件写入失败")
		}
	}
	return nil
}

func AdvancedWriteFile() error {

	path := path.Join(basePath, "demo3.json")
	err := os.WriteFile(path, []byte("hello,world\nhello,world\n"), 0666)
	if err != nil {
		return err
	}

	return nil
}

func IOWriteFile() error {
	file, err := os.OpenFile(newFilePath, os.O_RDWR|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	fmt.Println("文件打开成功： ", file.Name())
	for i := 0; i < 6; i++ {
		n, err := io.WriteString(file, "Hello,world\n")
		if err != nil {
			return errors.Wrap(err, "文件写入失败")
		}
		fmt.Println(n)
	}
	return nil

}
