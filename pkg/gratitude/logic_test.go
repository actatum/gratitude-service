package gratitude

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSendPrivate(t *testing.T) {
	tt := []struct {
		name     string
		req      *Message
		expected *SendResponse
	}{
		{name: "send private", req: &Message{
				Sender: User{
					UID:      "y",
					Name:     "y",
					Email:    "y@y.y",
					ImageURL: "",
				},
				Receivers: []User{{
					UID:      "x",
					Name:     "you",
					Email:    "you@you.you",
					ImageURL: "",
				}},
				Text:   "message",
				SentAt: time.Time{},
				ID:     "1",
				Kind:   "private",
		}, expected: &SendResponse{
			ID: "1",
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
			assert.Equal(t, tc.expected.ID, res.ID)
		})
	}
}

func TestSendPublic(t *testing.T) {
	tt := []struct {
		name     string
		req      *Message
		expected *SendResponse
	}{
		{name: "send public", req: &Message{
			Sender: User{
				UID:      "y",
				Name:     "y",
				Email:    "y@y.y",
				ImageURL: "",
			},
			Receivers: []User{{
				UID:      "x",
				Name:     "you",
				Email:    "you@you.you",
				ImageURL: "",
			}},
			Text:   "message",
			SentAt: time.Time{},
			ID:     "1",
			Kind:   "private",
		}, expected: &SendResponse{
			ID: "1",
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
			assert.Equal(t, tc.expected.ID, res.ID)
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
			Messages: []Message{{
				Sender: User{
					UID:      "y",
					Name:     "y",
					Email:    "y@y.y",
					ImageURL: "",
				},
				Receivers: []User{{
					UID:      "x",
					Name:     "you",
					Email:    "you@you.you",
					ImageURL: "",
				}},
				Text:   "message",
				SentAt: time.Time{},
				ID:     "1",
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
			assert.Equal(t, tc.expected.Messages[0].Receivers, res.Messages[0].Receivers)
			assert.Equal(t, tc.expected.Messages[0].Text, res.Messages[0].Text)
			assert.Equal(t, tc.expected.Messages[0].ID, res.Messages[0].ID)
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
			Messages: []Message{{
				Sender: User{
					UID:      "y",
					Name:     "y",
					Email:    "y@y.y",
					ImageURL: "",
				},
				Receivers: []User{{
					UID:      "x",
					Name:     "you",
					Email:    "you@you.you",
					ImageURL: "",
				}},
				Text:   "message",
				SentAt: time.Time{},
				ID:     "1",
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
			assert.Equal(t, tc.expected.Messages[0].Receivers, res.Messages[0].Receivers)
			assert.Equal(t, tc.expected.Messages[0].Text, res.Messages[0].Text)
			assert.Equal(t, tc.expected.Messages[0].ID, res.Messages[0].ID)
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
			Messages: []Message{{
				Sender: User{
					UID:      "y",
					Name:     "y",
					Email:    "y@y.y",
					ImageURL: "",
				},
				Receivers: []User{{
					UID:      "x",
					Name:     "you",
					Email:    "you@you.you",
					ImageURL: "",
				}},
				Text:   "message",
				SentAt: time.Time{},
				ID:     "1",
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
			assert.Equal(t, tc.expected.Messages[0].Receivers, res.Messages[0].Receivers)
			assert.Equal(t, tc.expected.Messages[0].Text, res.Messages[0].Text)
			assert.Equal(t, tc.expected.Messages[0].ID, res.Messages[0].ID)
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
			Users: []User{{
				UID:      "1",
				Name:     "me",
				ImageURL: "http://image.com",
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
			assert.Equal(t, tc.expected.Users[0].UID, res.Users[0].UID)
			assert.Equal(t, tc.expected.Users[0].Name, res.Users[0].Name)
			assert.Equal(t, tc.expected.Users[0].ImageURL, res.Users[0].ImageURL)
		})
	}
}
