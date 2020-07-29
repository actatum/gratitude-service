package gratitude

import "context"

// Provider interface for identity provider functionality
type Provider interface {
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error)
}
