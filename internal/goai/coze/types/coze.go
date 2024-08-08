package types

type ChatReq struct {
	BotID              string               `json:"bot_id"`
	UserID             string               `json:"user_id"`
	Stream             bool                 `json:"stream"`
	AutoSaveHistory    bool                 `json:"auto_save_history"`
	AdditionalMessages []AdditionalMessages `json:"additional_messages"`
}

type AdditionalMessages struct {
	Role        string `json:"role"`
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
}

type ChatResp struct {
	Data Data   `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type LastError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Data struct {
	ID             string    `json:"id"`
	ConversationID string    `json:"conversation_id"`
	BotID          string    `json:"bot_id"`
	CreatedAt      int       `json:"created_at"`
	LastError      LastError `json:"last_error"`
	Status         string    `json:"status"`
}

type ChatMessage struct {
	Code int           `json:"code"`
	Data []MessageData `json:"data"`
	Msg  string        `json:"msg"`
}
type MessageData struct {
	BotID          string `json:"bot_id"`
	Content        string `json:"content"`
	ContentType    string `json:"content_type"`
	ConversationID string `json:"conversation_id"`
	ID             string `json:"id"`
	Role           string `json:"role"`
	Type           string `json:"type"`
}

type FlowMessageData struct {
	ID             string `json:"id"`
	ConversationID string `json:"conversation_id"`
	BotID          string `json:"bot_id"`
	Role           string `json:"role"`
	Type           string `json:"type"`
	Content        string `json:"content"`
	ContentType    string `json:"content_type"`
	ChatID         string `json:"chat_id"`
}
