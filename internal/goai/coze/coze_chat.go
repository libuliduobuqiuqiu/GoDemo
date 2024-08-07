package coze

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"godemo/internal/goai/coze/types"
	"net/http"
	"time"
)

const (
	CompletedStatus = "completed"
	ChatUrl         = "https://api.coze.cn/v3/chat"
	ChatDetailUrl   = "https://api.coze.cn/v3/chat/retrieve"
	ChatMessageUrl  = "https://api.coze.cn/v3/chat/message/list"
)

func GetChatDetail(ctx context.Context, client *http.Client, chat_id, conversation_id string) {
	url := fmt.Sprintf("%s?conversation_id=%s&chat_id=%s", ChatDetailUrl, conversation_id, chat_id)

loop:
	for {
		resp, err := Request(ctx, client, http.MethodGet, url, nil)
		if err != nil {
			log.WithError(err).Error()
			return
		}

		var chatResp types.ChatResp
		if err := json.Unmarshal(resp, &chatResp); err != nil {
			log.WithError(err).Error()
			return
		}

		if chatResp.Data.Status == "completed" {
			break loop
		}

		time.Sleep(1 * time.Second)
	}
}

func GetChatMessage(ctx context.Context, client *http.Client, chat_id, conversation_id string) {

	url := fmt.Sprintf("%s?conversation_id=%s&chat_id=%s", ChatMessageUrl, conversation_id, chat_id)

	resp, err := Request(ctx, client, http.MethodGet, url, nil)
	if err != nil {
		log.WithError(err).Error()
		return
	}

	var chatMessage types.ChatMessage
	if err := json.Unmarshal(resp, &chatMessage); err != nil {
		log.WithError(err).Error()
		return
	}

	for _, msg := range chatMessage.Data {
		if msg.Type != "verbose" {
			fmt.Println(msg.Content)
		}
	}

}
