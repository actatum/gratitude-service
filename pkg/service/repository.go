package gratitude

import "context"

// Repository interface for interactions with databases
type Repository interface {
	SendPrivate(context.Context, *Message) (*MessageResponse, error)
	SendPublic(context.Context, *Message) (*MessageResponse, error)
	GetAllPublic(context.Context, *GetAllPublicRequest) (*GetAllPublicResponse, error)
	GetAllInbox(context.Context, *GetAllInboxRequest) (*GetAllInboxResponse, error)
	GetAllOutbox(context.Context, *GetAllOutboxRequest) (*GetAllOutboxResponse, error)
}
