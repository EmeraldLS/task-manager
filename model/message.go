package model

// Message represents a simple message with a body.
type Message struct {
	Body string
}

// NewMessage creates a new Message with the given body.
func NewMessage(body string) *Message {
	return &Message{
		Body: body,
	}
}
