//go:build integration

package gotool

import (
	"fmt"
	"log"
	"testing"

	"godemo/internal/gotool/cryptodemo"
)

func TestParseCert(t *testing.T) {
	filePath := "/mnt/d/Company/regualr_cert.p12"

	fp, err := cryptodemo.FingerprintP12SHA256(filePath, "123456")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fp)
}
