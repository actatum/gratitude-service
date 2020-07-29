package api

import (
	"net/http"

	"github.com/actatum/gratitude-board-service/pkg/gratitude"
)

func validateMessage(msg *gratitude.Message) error {
	if msg.Sender == "" {
		return gratitude.NewGratitudeError(http.StatusBadRequest, "invalid request: message sender is empty")
	}
	if msg.SenderID == "" {
		return gratitude.NewGratitudeError(http.StatusBadRequest, "invalid request: message sender ID is empty")
	}
	if msg.Recipients == nil {
		return gratitude.NewGratitudeError(http.StatusBadRequest, "invalid request: message recipient is empty")
	}
	if msg.RecipientIDs == nil {
		return gratitude.NewGratitudeError(http.StatusBadRequest, "invalid request: message recipient ID is empty")
	}
	if msg.Text == "" {
		return gratitude.NewGratitudeError(http.StatusBadRequest, "invalid request: message text is empty")
	}

	return nil
}
