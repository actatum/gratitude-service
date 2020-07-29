package gratitude

import (
	"context"

	errs "github.com/pkg/errors"
)

type service struct {
	provider Provider
	repo     Repository
}

// NewGratitudeService returns an object implementing the GratitudeService interface
func NewGratitudeService(provider Provider, repository Repository) GratitudeService {
	return &service{
		provider: provider,
		repo:     repository,
	}
}

func (s *service) SendPrivate(ctx context.Context, req *Message) (*SendResponse, error) {
	res, err := s.repo.SendPrivate(ctx, req)
	if err != nil {
		return nil, errs.Wrap(err, "service.Gratitude.SendPrivate")
	}

	return res, nil
}

func (s *service) SendPublic(ctx context.Context, req *Message) (*SendResponse, error) {
	res, err := s.repo.SendPublic(ctx, req)
	if err != nil {
		return nil, errs.Wrap(err, "service.Gratitude.SendPublic")
	}

	return res, nil
}

func (s *service) GetAllPublic(ctx context.Context, req *GetAllPublicRequest) (*GetAllPublicResponse, error) {
	res, err := s.repo.GetAllPublic(ctx, req)
	if err != nil {
		return nil, errs.Wrap(err, "service.Gratitude.GetAllPublic")
	}

	return res, nil
}

func (s *service) GetAllInbox(ctx context.Context, req *GetAllInboxRequest) (*GetAllInboxResponse, error) {
	res, err := s.repo.GetAllInbox(ctx, req)
	if err != nil {
		return nil, errs.Wrap(err, "service.Gratitude.GetAllInbox")
	}

	return res, nil
}

func (s *service) GetAllOutbox(ctx context.Context, req *GetAllOutboxRequest) (*GetAllOutboxResponse, error) {
	res, err := s.repo.GetAllOutbox(ctx, req)
	if err != nil {
		return nil, errs.Wrap(err, "service.Gratitude.GetAllOutbox")
	}

	return res, nil
}

func (s *service) GetAllUsers(ctx context.Context, req *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	res, err := s.provider.GetAllUsers(ctx, req)
	if err != nil {
		return nil, errs.Wrap(err, "service.Gratitude.GetAllUsers")
	}

	return res, nil
}
