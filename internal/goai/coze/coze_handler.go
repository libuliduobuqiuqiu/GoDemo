package coze

import (
	"bufio"
	"context"
	"fmt"
	"godemo/internal/goai/coze/types"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	productName = "coze"
)

func StartCozeChat() {
	client := GenCozeClient()

	ctx := context.Background()
	cozeClient := &CozeClient{Client: client, IsStream: true}

	var additionalMessages []types.AdditionalMessages
	var url = ChatUrl
	req := types.ChatReq{
		BotID:           BotID,
		UserID:          UserID,
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
		additionalMessages = []types.AdditionalMessages{msg}
		req.AdditionalMessages = additionalMessages

		if cozeClient.ConversationID != "" {
			url += fmt.Sprintf("?conversation_id=%s", cozeClient.ConversationID)
		}

		fmt.Println("KayChat:")
		_, err := cozeClient.Request(ctx, http.MethodPost, url, req)
		fmt.Println(additionalMessages)
		if err != nil {
			log.WithError(err).Error()
		}

		cozeClient.GetChatMessage(ctx, cozeClient.ChatID, cozeClient.ConversationID)

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
