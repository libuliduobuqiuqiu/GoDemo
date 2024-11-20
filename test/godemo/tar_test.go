package godemo

import (
	"godemo/internal/godemo"
	"testing"
)

func TestExtractTGZ(t *testing.T) {
	err := godemo.ExtractTGZ("/data/Bak/ebook.tgz", "/data/Bak/")
	if err != nil {
		t.Fatal(err)
	}
}
