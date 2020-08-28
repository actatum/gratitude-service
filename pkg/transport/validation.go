package transport

import (
	"net/http"

	"github.com/actatum/gratitude-board-service/pkg/gratitude"
)

func validateMessage(msg *gratitude.Message) error {
	if msg.Sender.Email == "" {
		return gratitude.NewGratitudeError(http.StatusBadRequest, "invalid request: message sender is empty")
	}

	if msg.Sender.UID == "" {
		return gratitude.NewGratitudeError(http.StatusBadRequest, "invalid request: message sender ID is empty")
	}

	if msg.Receivers == nil {
		return gratitude.NewGratitudeError(http.StatusBadRequest, "invalid request: message recipient is empty")
	}

	if msg.Text == "" {
		return gratitude.NewGratitudeError(http.StatusBadRequest, "invalid request: message text is empty")
	}

	return nil
}
