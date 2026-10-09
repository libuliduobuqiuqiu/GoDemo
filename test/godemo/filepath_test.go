package godemo

import (
	"fmt"
	"godemo/internal/godemo"
	"sort"
	"testing"
)

func TestFilePathAbs(t *testing.T) {

	godemo.GetAbsPath()

}

func TestScanDir(t *testing.T) {
	dirPath := "."

	files, err := godemo.ScanDir(dirPath)
	if err != nil {
		t.Fatal(err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].ModifiedTime.After(files[j].ModifiedTime)
	})

	for _, f := range files {
		fmt.Printf("%+v\n", f)
	}

	var tmpErr error

	fmt.Println(tmpErr == nil)
	fmt.Println(&tmpErr == nil)

}
