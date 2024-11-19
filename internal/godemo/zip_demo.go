package godemo

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
