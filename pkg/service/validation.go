package gratitude

import "fmt"

type messageError struct {
	err   string
	field string
}

func (e *messageError) Error() string {
	return fmt.Sprintf("invalid request: message %s is empty", e.field)
}

func validateMessage(msg *Message) error {
	if msg.Sender == "" {
		return &messageError{err: "", field: "sender"}
	}
	if msg.SenderId == "" {
		return &messageError{err: "", field: "sender ID"}
	}
	if msg.Recipients == nil {
		return &messageError{err: "", field: "recipient"}
	}
	if msg.RecipientIds == nil {
		return &messageError{err: "", field: "recipient ID"}
	}
	if msg.Text == "" {
		return &messageError{err: "", field: "text"}
	}

	return nil
}
