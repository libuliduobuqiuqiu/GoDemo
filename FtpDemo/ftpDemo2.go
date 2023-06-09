package main

import (
	"errors"
	"fmt"
	"github.com/jlaffaye/ftp"
	"runtime"
	"sync"
	"time"
)

const (
	Data_Size = 1024
)

// FTP 文件对象
type FileSource struct {
	entry *ftp.Entry // ftp库的entry对象
	path  string     // 文件的全路径+文件名
}

// EntryHandler 遍历ftp目录时的文件handler
type EntryHandler func(e *ftp.Entry, currentPath string) error

// FTP文件信息
type FtpFile struct {
	FileName string //FTP文件名
	Path     string //FTP文件的全路径+文件名
	Type     int    //FTP文件类型，文件:0, 文件夹:1
	Size     int    //FTP文件大小
}
type Deal struct {
	ftp      *ftp.ServerConn
	fileChan chan interface{}
	wg       sync.WaitGroup
}

func NewDeal() *Deal {
	return &Deal{}
}
func (this *Deal) Init(addr, user, passwd string) error {
	var err error
	this.ftp, err = ftp.Connect(addr)
	if err != nil {
		return err
	}
	err = this.ftp.Login(user, passwd)
	if err != nil {
		return err
	}
	this.fileChan = make(chan interface{}, Data_Size)
	this.wg = sync.WaitGroup{}
	fmt.Println("ftp连接成功")
	return nil
}
func (this *Deal) Fini() error {
	if this.ftp == nil {
		return errors.New("FTP客户端指针为空,注销失败")
	}
	err := this.ftp.Logout()
	if err != nil {
		fmt.Println("FTP注销失败,error info:", err)
		return err
	}
	return nil
}
func (this *Deal) Process(addr, user, passwd, rootDir string) error {
	err := this.Init(addr, user, passwd)
	if err != nil {
		fmt.Println("初始化失败")
		return err
	}
	for {
		//第一种 在递归目录中获取文件列表，耗时最小
		this.walk(rootDir)
		//第二种 在递归目录中回调函数中获取文件列表,多级目录耗时大
		//this.listfiles(rootDir)
		//第三种 在递归目录中回调函数中获取文件列表，多级目录耗时大
		//this.walkCall(rootDir, this.Handler)
		go func() {
			defer this.wg.Done()
			for {
				select {
				case data := <-this.fileChan:
					stru := data.(FtpFile)
					fmt.Println("文件名:", stru.FileName)
					break
				case <-time.After(time.Millisecond * 100):
					runtime.Gosched() //切换任务
					break
				}
			}
		}()
		this.wg.Wait()
		time.Sleep(time.Second)
	}
	return nil
}

// 函调函数
func (this *Deal) Handler(e *ftp.Entry, currentPath string) error {
	stru := FtpFile{}
	stru.FileName = e.Name
	stru.Path = currentPath + "//" + e.Name //CKK/20191102/10/17/20191113170257659_1d1d19f4-dd2f-4662-af8c-30658bd1e90.zlib
	stru.Type = int(e.Type)
	stru.Size = int(e.Size)
	select {
	case this.fileChan <- stru:
		//global.Log.Debug("fileChandata: %v", stru)
	default:
		fmt.Println("fileChan data chan is full")
		time.Sleep(time.Second)
		break
	}
	return nil
}

// 遍历ftp目录，获取文件
func (this *Deal) walk(rootDir string) error {
	entries, err := this.ftp.List(rootDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		switch entry.Type {
		case ftp.EntryTypeFile:
			//正在上传的文件，先不进行下载
			if entry.Size != 0 {
				stru := FtpFile{}
				stru.FileName = entry.Name
				stru.Path = rootDir + "//" + entry.Name //CKK/20191102/10/17/20191113170257659_1d1d19f4-dd2f-4662-af8c-30658bd1e90.zlib
				stru.Type = int(entry.Type)
				stru.Size = int(entry.Size)
				if len(this.fileChan) > (Data_Size - 1) {
					fmt.Println("管道大小超限，完成本地扫描:", len(this.fileChan))
					return nil
				}
				select {
				case this.fileChan <- stru:
					//global.Log.Debug("fileChan: %v", stru.FileName)
				default:
					fmt.Println("fileChan data chan is full")
					time.Sleep(time.Second)
					return nil
				}
			}
		case ftp.EntryTypeFolder:
			this.walk(fmt.Sprintf("%s/%s", rootDir, entry.Name))
		default:
		}
	}
	return nil
}

// 遍历ftp目录，回调获取文件
func (this *Deal) walkCall(rootDir string, handler EntryHandler) error {
	entries, err := this.ftp.List(rootDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		switch entry.Type {
		case ftp.EntryTypeFile:
			//正在上传的文件，先不进行下载
			if entry.Size != 0 {
				handler(entry, rootDir)
			}
		case ftp.EntryTypeFolder:
			this.walkCall(fmt.Sprintf("%s/%s", rootDir, entry.Name), handler)
		default:
		}
	}
	return nil
}

// 遍历ftp目录，获取文件
func (this *Deal) listfiles(rootDir string) error {
	err := this.walkCall(rootDir, func(entry *ftp.Entry, currentPath string) error {
		stru := FtpFile{}
		stru.FileName = entry.Name
		stru.Path = currentPath + "//" + entry.Name //CKK/20191102/10/17/20191113170257659_1d1d19f4-dd2f-4662-af8c-30658bd1e90.zlib
		stru.Type = int(entry.Type)
		stru.Size = int(entry.Size)
		if len(this.fileChan) > Data_Size {
			return nil
		}
		select {
		case this.fileChan <- stru:
			fmt.Println("fileChan:", stru.FileName)
		default:
			fmt.Println("fileChan data chan is full")
			time.Sleep(time.Second)
			break
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
