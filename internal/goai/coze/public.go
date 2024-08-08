package coze

import (
	"crypto/tls"
	"godemo/pkg"
	"net/http"
)

var (
	Authorization string
	UserID        string
	BotID         string
)

func init() {
	InitAuthorization()
}

func GenCozeClient() *http.Client {
	transport := http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: &transport}
}

func InitAuthorization() {
	data := pkg.GetAISecuretConfig("coze")

	if config, ok := data.(map[string]interface{}); ok {
		for k := range config {
			switch k {
			case "authorization":
				Authorization = config["authorization"].(string)
			case "user_id":
				UserID = config["user_id"].(string)
			case "bot_id":
				BotID = config["bot_id"].(string)
			}
		}
	}
}
