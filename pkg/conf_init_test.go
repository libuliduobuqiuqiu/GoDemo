package pkg

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetGlobalConfig(t *testing.T) {
	t.Setenv(ConfigPathEnv, "")
	if _, err := GetGlobalConfig(""); err == nil || !strings.Contains(err.Error(), ConfigPathEnv) {
		t.Fatalf("missing config path: got %v", err)
	}

	dir := t.TempDir()
	validPath := filepath.Join(dir, "config.json")
	data := `{"mysql":{"username":"demo","password":"secret","host":"localhost","port":3306,"prefix":"sample"},"devices":[{"host":"127.0.0.1","port":22}]}`
	if err := os.WriteFile(validPath, []byte(data), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv(ConfigPathEnv, filepath.Join(dir, "missing.json"))
	config, err := GetGlobalConfig(validPath)
	if err != nil {
		t.Fatal(err)
	}
	if config.MysqlConfig.Username != "demo" || config.MysqlConfig.Prefix != "sample" || len(config.Devices) != 1 {
		t.Fatalf("unexpected config: %+v", config)
	}

	t.Setenv(ConfigPathEnv, validPath)
	if _, err := GetGlobalConfig(""); err != nil {
		t.Fatalf("environment path: %v", err)
	}

	missingPath := filepath.Join(dir, "missing.json")
	if _, err := GetGlobalConfig(missingPath); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("missing file: got %v", err)
	}

	badPath := filepath.Join(dir, "bad.json")
	if err := os.WriteFile(badPath, []byte("{"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := GetGlobalConfig(badPath); err == nil || !strings.Contains(err.Error(), badPath) {
		t.Fatalf("invalid JSON: got %v", err)
	}
}

func TestGenMysqlDSN(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(path, []byte(`{"mysql":{"username":"demo","password":"secret","host":"localhost","port":3306,"prefix":"sample"}}`), 0600); err != nil {
		t.Fatal(err)
	}
	dsn, err := GenMysqlDSN(path)
	if err != nil {
		t.Fatal(err)
	}
	if want := "demo:secret@tcp(localhost:3306)/sample?charset=utf8mb4&parseTime=True"; dsn != want {
		t.Fatalf("DSN = %q, want %q", dsn, want)
	}
}
