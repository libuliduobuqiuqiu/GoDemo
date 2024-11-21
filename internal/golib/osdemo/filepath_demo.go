package osdemo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
		return err
	}
	return nil
}
