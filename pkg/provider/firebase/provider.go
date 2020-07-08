package firebase

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	gratitude "github.com/actatum/gratitude-board-service/pkg/service"
	errs "github.com/pkg/errors"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var (
	projectID = os.Getenv("PROJECT_ID")
)

type provider struct {
	client *auth.Client
}

func newFirebaseAuthClient() (*auth.Client, error) {
	ctx := context.Background()
	creds := os.Getenv("GCP")
	opt := option.WithCredentialsJSON([]byte(creds))
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, errs.Wrap(err, "provider.Firebase.newFirebaseAuthClient")
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "provider.Firebase.newFirebaseAuthClient")
	}

	return client, nil
}

// NewFirebaseProvider returns an implementation of the provider interface using google's Firebase
func NewFirebaseProvider() (gratitude.Provider, error) {
	provider := &provider{}
	client, err := newFirebaseAuthClient()
	if err != nil {
		return nil, errs.Wrap(err, "repository.Firebase.NewFirebaseProvider")
	}

	provider.client = client

	return provider, nil
}

func (p *provider) GetAllUsers(ctx context.Context, req *gratitude.GetAllUsersRequest) (*gratitude.GetAllUsersResponse, error) {
	var res gratitude.GetAllUsersResponse
	iter := p.client.Users(ctx, "")
	for {
		user, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errs.Wrap(err, "provider.Firebase.GetAllUsers")
		}
		u := &gratitude.User{
			Uid:      user.UID,
			Name:     user.DisplayName,
			Email:    user.Email,
			ImageUrl: user.PhotoURL,
		}
		res.Users = append(res.Users, u)
	}

	return &res, nil
}
