package coze

import (
	"encoding/json"
	"fmt"
	"godemo/internal/goai/coze/types"
	"strings"
)

func (c *CozeClient) Write(body []byte) (n int, err error) {
	c.ConvertCozeRespToContent(body)
	return len(body), nil
}

func (c *CozeClient) ConvertCozeRespToContent(body []byte) (completed bool) {
	resp := string(body)

	if resp == "\n" {
		return
	}

	tmpResp := strings.Split(resp, ":")

	if len(tmpResp) > 2 {

		switch tmpResp[0] {
		case "event":
			if tmpResp[1] == types.DoneEvent {
				completed = true
			}

		case "data":

			flowData := &types.FlowMessageData{}
			data := strings.Join(tmpResp[1:], ":")
			if err := json.Unmarshal([]byte(data), flowData); err != nil {
				fmt.Println("接收流式响应异常：", err.Error())
			}

			switch flowData.Type {
			case "answer":
				fmt.Printf(flowData.Content)
			case "verbose":
				c.ConversationID = flowData.ConversationID
				c.ChatID = flowData.ChatID
				fmt.Printf("\n你可能还想知道的：\n")
			case "follow_up":
				fmt.Println(flowData.Content)
			}
		}

	}

	return
}
