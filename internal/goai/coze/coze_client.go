package coze

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"godemo/internal/goai/coze/types"
	"godemo/pkg"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	productName = "coze"
)

var (
	authorization string
	userID        string
	botID         string
)

func genCozeClient() *http.Client {
	transport := http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: &transport}
}

func initAuthorization() {
	data := pkg.GetAISecuretConfig("coze")

	if config, ok := data.(map[string]interface{}); ok {
		for k := range config {
			switch k {
			case "authorization":
				authorization = config["authorization"].(string)
			case "user_id":
				userID = config["user_id"].(string)
			case "bot_id":
				botID = config["bot_id"].(string)
			}
		}
	}
}

func Request(ctx context.Context, client *http.Client, method, url string, reqData interface{}) (data []byte, err error) {
	var (
		body io.Reader
	)

	// 序列化
	if reqData != nil {
		tmpData, err := json.Marshal(reqData)
		if err != nil {
			log.WithError(err).Error()
		}
		body = bytes.NewReader(tmpData)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func UseCozeChat() {
	initAuthorization()
	client := genCozeClient()
	ctx := context.Background()

	var additionalMessages []types.AdditionalMessages
	req := types.ChatReq{
		BotID:           botID,
		UserID:          userID,
		Stream:          true,
		AutoSaveHistory: true,
	}

	fmt.Println("Welcome use kay chat!")
	buffer := bufio.NewScanner(os.Stdin)

	for buffer.Scan() {
		var text = buffer.Text()

		msg := types.AdditionalMessages{
			Role:        "user",
			ContentType: "text",
			Content:     text,
		}
		additionalMessages = append(additionalMessages, msg)
		req.AdditionalMessages = additionalMessages

		go func() {
			resp, err := Request(ctx, client, http.MethodPost, ChatUrl, req)
			if err != nil {
				log.WithError(err).Error()
			}

			fmt.Println(string(resp))
		}()

		// var chatResp types.ChatResp
		// if err := json.Unmarshal(resp, &chatResp); err != nil {
		// 	log.WithError(err).Error()
		// 	return
		// }

		// if chatResp.Data.Status != CompletedStatus {
		// 	GetChatDetail(ctx, client, chatResp.Data.ID, chatResp.Data.ConversationID)
		// 	fmt.Println("Kay Chat:")
		// 	GetChatMessage(ctx, client, chatResp.Data.ID, chatResp.Data.ConversationID)
		// }
	}
}
