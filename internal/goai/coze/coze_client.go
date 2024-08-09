package coze

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"godemo/internal/goai/coze/types"
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	CompletedStatus = "completed"
	ChatUrl         = "https://api.coze.cn/v3/chat"
	ChatDetailUrl   = "https://api.coze.cn/v3/chat/retrieve"
	ChatMessageUrl  = "https://api.coze.cn/v3/chat/message/list"
)

type CozeClient struct {
	*http.Client
	IsStream       bool
	ConversationID string
	ChatID         string
}

func (c *CozeClient) Request(ctx context.Context, method, url string, reqData interface{}) (data []byte, err error) {
	var (
		body io.Reader
	)

	// 序列化请求
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
	req.Header.Add("Authorization", "Bearer "+Authorization)

	resp, err := c.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	// 读取响应内容
	if c.IsStream {
		if method == http.MethodPost {

			scanner := bufio.NewScanner(resp.Body)
			for scanner.Scan() {
				fmt.Fprintln(c, scanner.Text())
			}

			return
		}
	}

	fmt.Println("End.")

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func (c *CozeClient) GetChatDetail(ctx context.Context, chat_id, conversation_id string) {
	fmt.Println(chat_id, conversation_id)
	url := fmt.Sprintf("%s?conversation_id=%s&chat_id=%s", ChatDetailUrl, conversation_id, chat_id)

loop:
	for {
		resp, err := c.Request(ctx, http.MethodGet, url, nil)
		if err != nil {
			log.WithError(err).Error()
			return
		}

		fmt.Println(string(resp))
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

func (c *CozeClient) GetChatMessage(ctx context.Context, chat_id, conversation_id string) {

	url := fmt.Sprintf("%s?conversation_id=%s&chat_id=%s", ChatMessageUrl, conversation_id, chat_id)

	resp, err := c.Request(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.WithError(err).Error()
		return
	}
	fmt.Println(string(resp))

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
