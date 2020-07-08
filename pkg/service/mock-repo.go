package gratitude

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) SendPrivate(ctx context.Context, req *Message) (*MessageResponse, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*MessageResponse), args.Error(1)
}

func (m *mockRepo) SendPublic(ctx context.Context, req *Message) (*MessageResponse, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*MessageResponse), args.Error(1)
}

func (m *mockRepo) GetAllPublic(ctx context.Context, req *GetAllPublicRequest) (*GetAllPublicResponse, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*GetAllPublicResponse), args.Error(1)
}

func (m *mockRepo) GetAllInbox(ctx context.Context, req *GetAllInboxRequest) (*GetAllInboxResponse, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*GetAllInboxResponse), args.Error(1)
}

func (m *mockRepo) GetAllOutbox(ctx context.Context, req *GetAllOutboxRequest) (*GetAllOutboxResponse, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*GetAllOutboxResponse), args.Error(1)
}
