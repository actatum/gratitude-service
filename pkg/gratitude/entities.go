package gratitude

import "time"

// User data model
type User struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageURL string `json:"image_url"`
}

// Message is a data model for a
type Message struct {
	ID            string    `json:"id"`
	Kind          string    `json:"kind"`
	SenderID      string    `json:"sender_id"`
	Sender        string    `json:"sender"`
	Recipients    []string  `json:"recipients"`
	RecipientsIDs []string  `json:"recipients_ids`
	Text          string    `json:"text"`
	SentAt        time.Time `json:"sent_at"`
	Seen          bool      `json:"seen"`
}

// SendResponse is a response model for sending public or private messages
type SendResponse struct {
	ID string `json:"id"`
}

// GetAllPublicRequest is a request model for retrieving all public messages
type GetAllPublicRequest struct{}

// GetAllPublicResponse is a response model for retrieving all public messages
type GetAllPublicResponse struct {
	Messages []Message `json:"messages"`
}

// GetAllInboxRequest is a request model for retrieving all inbox messages
type GetAllInboxRequest struct{}

// GetAllInboxResponse is a response model for retrieving all inbox messages
type GetAllInboxResponse struct {
	Messages []Message `json:"messages"`
}

// GetAllOutboxRequest is a request model for retrieving all outbox messages
type GetAllOutboxRequest struct{}

// GetAllOutboxResponse is a response model for retrieving all outbox messages
type GetAllOutboxResponse struct {
	Messages []Message `json:"messages"`
}

// GetAllUsersRequest is a request model for retrieving all users on the platform
type GetAllUsersRequest struct{}

// GetAllUsersResponse is a response model for retrieving all users on the platform
type GetAllUsersResponse struct {
	Users []User `json:"users"`
}
