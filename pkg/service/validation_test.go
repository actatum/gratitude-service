package gratitude

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validateMessage(t *testing.T) {
	tt := []struct {
		name    string
		msg     *Message
		wantErr bool
		err     string
	}{
		{name: "no sender", msg: &Message{
			Sender: "",
		}, wantErr: true, err: "invalid request: message sender is empty"},
		{name: "no sender ID", msg: &Message{
			Sender:   "me",
			SenderId: "",
		}, wantErr: true, err: "invalid request: message sender ID is empty"},
		{name: "no recipients", msg: &Message{
			Sender:     "me",
			SenderId:   "id",
			Recipients: nil,
		}, wantErr: true, err: "invalid request: message recipient is empty"},
		{name: "no recipient IDs", msg: &Message{
			Sender:       "me",
			SenderId:     "id",
			Recipients:   []string{"you"},
			RecipientIds: nil,
		}, wantErr: true, err: "invalid request: message recipient ID is empty"},
		{name: "no text", msg: &Message{
			Sender:       "me",
			SenderId:     "id",
			Recipients:   []string{"you"},
			RecipientIds: []string{"id"},
			Text:         "",
		}, wantErr: true, err: "invalid request: message text is empty"},
		{name: "valid message", msg: &Message{
			Sender:       "me",
			SenderId:     "id",
			Recipients:   []string{"you"},
			RecipientIds: []string{"id"},
			Text:         "heres some text",
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
