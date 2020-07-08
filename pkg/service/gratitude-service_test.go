package gratitude

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendPrivate(t *testing.T) {
	tt := []struct {
		name     string
		req      *Message
		expected *MessageResponse
	}{
		{name: "send private", req: &Message{
			Sender:       "me",
			SenderId:     "id",
			Recipients:   []string{"you"},
			RecipientIds: []string{"id"},
			Text:         "message",
			SentAt:       nil,
			MessageId:    "1",
			Kind:         "private",
		}, expected: &MessageResponse{
			MessageId: "1",
		}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockRepo{}
			mockRepo.On("SendPrivate").Return(tc.expected, nil)

			testService := NewGratitudeService(nil, mockRepo)

			res, _ := testService.SendPrivate(context.Background(), tc.req)

			// Mock assertion: behavioral
			mockRepo.AssertExpectations(t)

			// Data assertion
			assert.Equal(t, tc.expected.MessageId, res.MessageId)
		})
	}
}

func TestSendPublic(t *testing.T) {
	tt := []struct {
		name     string
		req      *Message
		expected *MessageResponse
	}{
		{name: "send public", req: &Message{
			Sender:       "me",
			SenderId:     "id",
			Recipients:   []string{"you"},
			RecipientIds: []string{"id"},
			Text:         "message",
			SentAt:       nil,
			MessageId:    "1",
			Kind:         "public",
		}, expected: &MessageResponse{
			MessageId: "1",
		}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockRepo{}
			mockRepo.On("SendPublic").Return(tc.expected, nil)

			testService := NewGratitudeService(nil, mockRepo)

			res, _ := testService.SendPublic(context.Background(), tc.req)

			// Mock assertion: behavioral
			mockRepo.AssertExpectations(t)

			// Data assertion
			assert.Equal(t, tc.expected.MessageId, res.MessageId)
		})
	}
}

func TestGetAllPublic(t *testing.T) {
	tt := []struct {
		name     string
		req      *GetAllPublicRequest
		expected *GetAllPublicResponse
	}{
		{name: "get all public", req: &GetAllPublicRequest{}, expected: &GetAllPublicResponse{
			Messages: []*Message{{
				Sender:       "me",
				SenderId:     "id",
				Recipients:   []string{"you"},
				RecipientIds: []string{"id"},
				Text:         "message",
				SentAt:       nil,
				MessageId:    "1",
			}},
		}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockRepo{}
			mockRepo.On("GetAllPublic").Return(tc.expected, nil)

			testService := NewGratitudeService(nil, mockRepo)

			res, _ := testService.GetAllPublic(context.Background(), tc.req)

			// Mock assertion: behavioral
			mockRepo.AssertExpectations(t)

			// Data assertion
			assert.Equal(t, tc.expected.Messages[0].Sender, res.Messages[0].Sender)
			assert.Equal(t, tc.expected.Messages[0].SenderId, res.Messages[0].SenderId)
			assert.Equal(t, tc.expected.Messages[0].Recipients, res.Messages[0].Recipients)
			assert.Equal(t, tc.expected.Messages[0].RecipientIds, res.Messages[0].RecipientIds)
			assert.Equal(t, tc.expected.Messages[0].Text, res.Messages[0].Text)
			assert.Equal(t, tc.expected.Messages[0].MessageId, res.Messages[0].MessageId)
		})
	}
}

func TestGetAllInbox(t *testing.T) {
	tt := []struct {
		name     string
		req      *GetAllInboxRequest
		expected *GetAllInboxResponse
	}{
		{name: "get all inbox", req: &GetAllInboxRequest{}, expected: &GetAllInboxResponse{
			Messages: []*Message{{
				Sender:       "me",
				SenderId:     "id",
				Recipients:   []string{"you"},
				RecipientIds: []string{"id"},
				Text:         "message",
				SentAt:       nil,
				MessageId:    "1",
			}},
		}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockRepo{}
			mockRepo.On("GetAllInbox").Return(tc.expected, nil)

			testService := NewGratitudeService(nil, mockRepo)

			res, _ := testService.GetAllInbox(context.Background(), tc.req)

			// Mock assertion: behavioral
			mockRepo.AssertExpectations(t)

			// Data assertion
			assert.Equal(t, tc.expected.Messages[0].Sender, res.Messages[0].Sender)
			assert.Equal(t, tc.expected.Messages[0].SenderId, res.Messages[0].SenderId)
			assert.Equal(t, tc.expected.Messages[0].Recipients, res.Messages[0].Recipients)
			assert.Equal(t, tc.expected.Messages[0].RecipientIds, res.Messages[0].RecipientIds)
			assert.Equal(t, tc.expected.Messages[0].Text, res.Messages[0].Text)
			assert.Equal(t, tc.expected.Messages[0].MessageId, res.Messages[0].MessageId)
		})
	}
}

func TestGetAllOutbox(t *testing.T) {
	tt := []struct {
		name     string
		req      *GetAllOutboxRequest
		expected *GetAllOutboxResponse
	}{
		{name: "get all outbox", req: &GetAllOutboxRequest{}, expected: &GetAllOutboxResponse{
			Messages: []*Message{{
				Sender:       "me",
				SenderId:     "id",
				Recipients:   []string{"you"},
				RecipientIds: []string{"id"},
				Text:         "message",
				SentAt:       nil,
				MessageId:    "1",
			}},
		}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mockRepo{}
			mockRepo.On("GetAllOutbox").Return(tc.expected, nil)

			testService := NewGratitudeService(nil, mockRepo)

			res, _ := testService.GetAllOutbox(context.Background(), tc.req)

			// Mock assertion: behavioral
			mockRepo.AssertExpectations(t)

			// Data assertion
			assert.Equal(t, tc.expected.Messages[0].Sender, res.Messages[0].Sender)
			assert.Equal(t, tc.expected.Messages[0].SenderId, res.Messages[0].SenderId)
			assert.Equal(t, tc.expected.Messages[0].Recipients, res.Messages[0].Recipients)
			assert.Equal(t, tc.expected.Messages[0].RecipientIds, res.Messages[0].RecipientIds)
			assert.Equal(t, tc.expected.Messages[0].Text, res.Messages[0].Text)
			assert.Equal(t, tc.expected.Messages[0].MessageId, res.Messages[0].MessageId)
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	tt := []struct {
		name     string
		req      *GetAllUsersRequest
		expected *GetAllUsersResponse
	}{
		{name: "get all users", req: &GetAllUsersRequest{}, expected: &GetAllUsersResponse{
			Users: []*User{{
				Uid:      "1",
				Name:     "me",
				ImageUrl: "http://image.com",
			}},
		}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockProvider := &mockProvider{}
			mockProvider.On("GetAllUsers").Return(tc.expected, nil)

			testService := NewGratitudeService(mockProvider, nil)

			res, _ := testService.GetAllUsers(context.Background(), tc.req)

			// Mock assertion: behavioral
			mockProvider.AssertExpectations(t)

			// Data assertion
			assert.Equal(t, tc.expected.Users[0].Uid, res.Users[0].Uid)
			assert.Equal(t, tc.expected.Users[0].Name, res.Users[0].Name)
			assert.Equal(t, tc.expected.Users[0].ImageUrl, res.Users[0].ImageUrl)
		})
	}
}
