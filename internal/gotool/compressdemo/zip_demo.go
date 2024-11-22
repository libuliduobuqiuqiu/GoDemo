package compressdemo

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CompressFiles(dirName string, zipName string) error {

	info, err := os.Stat(dirName)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("Need dir path.")
	}

	zipFile, err := os.Create(zipName)
	if err != nil {
		return err
	}

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		realPath, err := filepath.Rel(dirName, path)

		if info.IsDir() {
			_, err = zipWriter.Create(realPath + "/")
			if err != nil {
				return err
			}
			return nil
		}

		fileWriter, err := zipWriter.Create(realPath)
		if err != nil {
			return err
		}

		sourceFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		_, err = io.Copy(fileWriter, sourceFile)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func UnCompressZip(zipName string, destDir string) (err error) {

	dirInfo, err := os.Stat(destDir)
	if err != nil {
		return
	}

	if !dirInfo.IsDir() {
		return fmt.Errorf("Need send a dir path.")
	}

	zip, err := zip.OpenReader(zipName)
	if err != nil {
		return
	}

	defer zip.Close()
	for _, file := range zip.File {
		path := filepath.Join(destDir, file.Name)
		info := file.FileInfo()

		if info.IsDir() {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}

		destFile, err := os.Create(path)
		if err != nil {
			return err
		}

		defer destFile.Close()

		srcFile, err := file.Open()
		if err != nil {
			return err
		}

		defer srcFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err
		}
	}

	return nil
}
