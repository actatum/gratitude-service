package transport

import (
	"testing"

	"github.com/actatum/gratitude-board-service/pkg/gratitude"
	"github.com/stretchr/testify/assert"
)

func Test_validateMessage(t *testing.T) {
	tt := []struct {
		name    string
		msg     *gratitude.Message
		wantErr bool
		err     string
	}{
		{name: "no sender", msg: &gratitude.Message{
			Sender: "",
		}, wantErr: true, err: "invalid request: message sender is empty"},
		{name: "no sender ID", msg: &gratitude.Message{
			Sender:   "me",
			SenderID: "",
		}, wantErr: true, err: "invalid request: message sender ID is empty"},
		{name: "no Receivers", msg: &gratitude.Message{
			Sender:    "me",
			SenderID:  "id",
			Receivers: nil,
		}, wantErr: true, err: "invalid request: message recipient is empty"},
		{name: "no text", msg: &gratitude.Message{
			Sender:   "me",
			SenderID: "id",
			Receivers: []gratitude.User{{
				UID:      "x",
				Name:     "you",
				Email:    "you@you.you",
				ImageURL: "",
			}},
			Text: "",
		}, wantErr: true, err: "invalid request: message text is empty"},
		{name: "valid message", msg: &gratitude.Message{
			Sender:   "me",
			SenderID: "id",
			Receivers: []gratitude.User{{
				UID:      "x",
				Name:     "you",
				Email:    "you@you.you",
				ImageURL: "",
			}},
			Text: "heres some text",
		}, wantErr: false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := validateMessage(tc.msg)

			if tc.wantErr {
				assert.NotNil(t, got)
				assert.Equal(t, tc.err, got.Error())
			} else {
				assert.Nil(t, got)
			}
		})
	}
}