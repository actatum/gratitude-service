package gratitude

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type mockProvider struct {
	mock.Mock
}

func (m *mockProvider) GetAllUsers(ctx context.Context, req *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*GetAllUsersResponse), args.Error(1)
}
