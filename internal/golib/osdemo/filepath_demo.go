package osdemo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

func PrintFilePath(dirName string) (err error) {

	fmt.Println(dirName)
	err = filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(dirName, path)
		if err != nil {
			return err
		}
		fmt.Println(relPath)

		afterPath := strings.TrimPrefix(relPath, string(filepath.Separator))
		fmt.Println(afterPath)

		return nil
	})

	if err != nil {
		logrus.WithError(err).Error()
		return err
	}
	return nil
}
