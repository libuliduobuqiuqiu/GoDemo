package pkg

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetAISecuretConfig(configName string) interface{} {
	var config = make(map[string]interface{})

	_, path, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to get file information.")
	}

	dir := filepath.Dir(path)
	parentDir := filepath.Dir(dir)
	confPath := filepath.Join(parentDir, "configs", "ai_conf.json")

	data, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}

	if v, ok := config[configName]; ok {
		return v
	}

	return nil
}
