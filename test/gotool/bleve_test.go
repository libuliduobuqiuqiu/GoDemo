//go:build integration

package gotool

import (
	"godemo/internal/gotool/blevedemo"
	"testing"
)

func TestWriteMessage(t *testing.T) {

	if err := blevedemo.WriteMessage(); err != nil {
		t.Fatal(err)
	}

}
