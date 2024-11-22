package compressdemo

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Tar 文件格式 [Header1][Content1][Header2][Content2]...[HeaderN][ContentN]
// 创建Tgz压缩文件
func CreateTGZ(srcDir string, tgzPath string) error {

	tgzFile, err := os.Create(tgzPath)
	if err != nil {
		return err
	}
	defer tgzFile.Close()

	gzWriter := gzip.NewWriter(tgzFile)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 计算相对路径
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		// 去除可能出现的前导分隔符..
		relPath = strings.TrimPrefix(relPath, string(filepath.Separator))
		fmt.Println(relPath)

		// 写入头信息
		header, err := tar.FileInfoHeader(info, path)
		if err != nil {
			return err
		}

		header.Name = relPath
		if err = tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// 如果是文件，保存文件信息
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(tarWriter, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// 解压Tgz压缩文件
func ExtractTGZ(tgzPath string, destDir string) error {
	tgzFile, err := os.Open(tgzPath)
	if err != nil {
		return err
	}

	gzipFile, err := gzip.NewReader(tgzFile)
	if err != nil {
		return err
	}
	defer gzipFile.Close()

	tarFile := tar.NewReader(gzipFile)
	for {
		header, err := tarFile.Next()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		targetPath := filepath.Join(destDir, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, os.ModePerm); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := writeFile(targetPath, tarFile); err != nil {
				return err
			}
		default:
			fmt.Printf("Ignore unsuport file type :%s\n", header.Name)
		}
	}
}

func writeFile(targetPath string, srcFile io.Reader) (err error) {

	if err = os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
		return
	}

	destFile, err := os.Create(targetPath)
	if err != nil {
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return
}
