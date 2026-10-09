//go:build integration

package goothers

import (
	"fmt"
	"godemo/internal/goothers"
	"os"
	"path"
	"testing"
)

func TestReplaceSnowFlake(t *testing.T) {
	dirPath := "/mnt/d/Company/tmp/replace_chains/"

	files, err := os.ReadDir(dirPath)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		filePath := path.Join(dirPath, f.Name())
		fmt.Println(filePath)
		if err = goothers.ReplaceUUIDWithSnowflakeID(filePath); err != nil {
			t.Fatal(err)
		}
	}
}
