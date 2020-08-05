package firebase

import (
	"context"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/actatum/gratitude-board-service/pkg/gratitude"
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
	creds := os.Getenv("GCP_KEY")
	opt := option.WithCredentialsJSON([]byte(creds))
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "provider.Firebase.newFirebaseAuthClient")
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "provider.Firebase.newFirebaseAuthClient")
	}

	return client, nil
}

// NewFirebaseProvider returns an implementation of the provider interface using google's Firebase
func NewFirebaseProvider() (gratitude.Provider, error) {
	provider := &provider{}
	client, err := newFirebaseAuthClient()
	if err != nil {
		return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firebase.NewFirebaseProvider")
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
			return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "provider.Firebase.GetAllUsers")
		}
		u := gratitude.User{
			UID:      user.UID,
			Name:     user.DisplayName,
			Email:    user.Email,
			ImageURL: user.PhotoURL,
		}
		res.Users = append(res.Users, u)
	}

	return &res, nil
}
