package types

const (
	ChatCreatedEvent        = "conversation.chat.created"
	ChatInProgressEvent     = "conversation.chat.in_progress"
	MessageDeltaEvent       = "conversation.message.delta"
	MessageCompletedEvent   = "conversation.message.completed"
	ChatCompletedEvent      = "conversation.chat.completed"
	ChatFailedEvent         = "conversation.chat.failed"
	ChatRequiresActionEvent = "conversation.chat.requires_action"
	ErrorEvent              = "error"
	DoneEvent               = "done"
)
