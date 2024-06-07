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

func IOReadFile(targetPath string) error {
	f, err := os.Open(targetPath)
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
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
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

// 简单读取文件然后写入另外一个文件
func SimpleCopyFile() error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = os.WriteFile(newFilePath, []byte(data), 0666)
	if err != nil {
		return err
	}

	return nil
}

func AdvancedCopyFile(originPath, targetPath string) error {
	origin, err := os.OpenFile(originPath, os.O_RDONLY, 0666)
	if err != nil {
		return errors.Wrap(err, "打开源文件失败")
	}
	defer origin.Close()

	target, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return errors.Wrap(err, "打开目的文件失败")
	}
	defer target.Close()

	n, err := target.ReadFrom(origin)
	if err != nil {
		return errors.Wrap(err, "复制文件失败")
	}
	fmt.Println("文件复制成功", n)
	return nil
}

func IOCopyFile(originPath, targetPath string) error {
	origin, err := os.OpenFile(originPath, os.O_RDONLY, 0666)
	if err != nil {
		return errors.Wrap(err, "打开源文件失败")
	}
	defer origin.Close()

	target, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return errors.Wrap(err, "打开目的文件失败")
	}
	defer target.Close()

	n, err := io.Copy(target, origin)
	if err != nil {
		return errors.Wrap(err, "复制文件失败")
	}
	fmt.Println("文件复制成功", n)
	return nil
}

func DeleteFile(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

func DeleteDir(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}
