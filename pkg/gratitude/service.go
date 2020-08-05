package gratitude

import "context"

// Service is the interface for the service layer
type Service interface {
	SendPrivate(context.Context, *Message) (*SendResponse, error)
	SendPublic(context.Context, *Message) (*SendResponse, error)
	GetAllPublic(context.Context, *GetAllPublicRequest) (*GetAllPublicResponse, error)
	GetAllInbox(context.Context, *GetAllInboxRequest) (*GetAllInboxResponse, error)
	GetAllOutbox(context.Context, *GetAllOutboxRequest) (*GetAllOutboxResponse, error)
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error)
}
