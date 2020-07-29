package gratitude

import "time"

type User struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageUrl string `json:"image_url"`
}

type Message struct {
	ID           string    `json:"id"`
	Kind         string    `json:"kind"`
	SenderID     string    `json:"sender_id"`
	Sender       string    `json:"sender"`
	Recipients   []string  `json:"recipients"`
	RecipientIDs []string  `json:"recipient_ids`
	Text         string    `json:"text"`
	SentAt       time.Time `json:"sent_at"`
}

type SendResponse struct {
	ID string `json:"id"`
}

type GetAllPublicRequest struct{}

type GetAllPublicResponse struct {
	Messages []Message `json:"messages"`
}

type GetAllInboxRequest struct{}

type GetAllInboxResponse struct {
	Messages []Message `json:"messages"`
}

type GetAllOutboxRequest struct{}

type GetAllOutboxResponse struct {
	Messages []Message `json:"messages"`
}

type GetAllUsersRequest struct{}

type GetAllUsersResponse struct {
	Users []User `json:"users"`
}
