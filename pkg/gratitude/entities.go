package gratitude

import "time"

// User data model
type User struct {
	UID      string `json:"uid" firestore:"uid"`
	Name     string `json:"name" firestore:"name"`
	Email    string `json:"email" firestore:"email"`
	ImageURL string `json:"image_url" firestore:"image_url"`
}

// Message is a data model for a
type Message struct {
	ID        string    `json:"id" firestore:"id"`
	Kind      string    `json:"kind" firestore:"kind"`
	Sender    User      `json:"sender" firestore:"sender"`
	Receivers []User    `json:"receivers" firestore:"receivers"`
	Text      string    `json:"text" firestore:"text"`
	SentAt    time.Time `json:"sent_at" firestore:"sent_at"`
	Seen      bool      `json:"seen" firestore:"seen"`
}

// SendResponse is a response model for sending public or private messages
type SendResponse struct {
	ID string `json:"id"`
}

// GetAllPublicRequest is a request model for retrieving all public messages
type GetAllPublicRequest struct {
	UID string `json:"uid"`
}

// GetAllPublicResponse is a response model for retrieving all public messages
type GetAllPublicResponse struct {
	Messages []Message `json:"messages"`
}

// GetAllInboxRequest is a request model for retrieving all inbox messages
type GetAllInboxRequest struct {
	UID string `json:"uid"`
}

// GetAllInboxResponse is a response model for retrieving all inbox messages
type GetAllInboxResponse struct {
	Messages []Message `json:"messages"`
}

// GetAllOutboxRequest is a request model for retrieving all outbox messages
type GetAllOutboxRequest struct {
	UID string `json:"uid"`
}

// GetAllOutboxResponse is a response model for retrieving all outbox messages
type GetAllOutboxResponse struct {
	Messages []Message `json:"messages"`
}

// GetAllUsersRequest is a request model for retrieving all users on the platform
type GetAllUsersRequest struct {
	UID string `json:"uid"`
}

// GetAllUsersResponse is a response model for retrieving all users on the platform
type GetAllUsersResponse struct {
	Users []User `json:"users"`
}
