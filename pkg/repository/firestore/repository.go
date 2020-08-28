package firestore

import (
	"context"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/actatum/gratitude-board-service/pkg/gratitude"
	errs "github.com/pkg/errors"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var (
	projectID = os.Getenv("PROJECT_ID")
)

type repository struct {
	client *firestore.Client
}

func newFirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()
	creds := os.Getenv("GCP_KEY")
	opt := option.WithCredentialsJSON([]byte(creds))
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firestore.newFirestoreClient")
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firestore.newFirestoreClient")
	}

	return client, nil
}

// NewFirestoreRepository creates a new firestore client and then
// returns an object satisfying the gratitude.Repository interface and an error
func NewFirestoreRepository() (gratitude.Repository, error) {
	repo := &repository{}
	client, err := newFirestoreClient()
	if err != nil {
		return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firestore.NewFirestoreRepository")
	}

	repo.client = client

	return repo, nil
}

func (r *repository) SendPrivate(ctx context.Context, req *gratitude.Message) (*gratitude.SendResponse, error) {
	var res gratitude.SendResponse

	// Add to senders outbox
	ref, _, err := r.client.Collection("users").Doc(req.Sender.UID).Collection("outbox").Add(ctx, req)
	if err != nil {
		return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firestore.SendPrivate")
	}

	// Add to all Receivers inbox
	for _, receiver := range req.Receivers {
		_, err = r.client.Collection("users").Doc(receiver.UID).Collection("inbox").Doc(ref.ID).Set(ctx, req)
		if err != nil {
			return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firestore.SendPublic")
		}
	}

	res.ID = ref.ID

	return &res, nil
}

func (r *repository) SendPublic(ctx context.Context, req *gratitude.Message) (*gratitude.SendResponse, error) {
	var res gratitude.SendResponse

	// Add to senders outbox
	ref, _, err := r.client.Collection("users").Doc(req.Sender.UID).Collection("outbox").Add(ctx, req)
	if err != nil {
		return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firestore.SendPublic")
	}

	// Add to all Receivers inbox
	for _, receiver := range req.Receivers {
		_, err = r.client.Collection("users").Doc(receiver.UID).Collection("inbox").Doc(ref.ID).Set(ctx, req)
		if err != nil {
			return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firestore.SendPublic")
		}
	}

	// Add to public message collection
	_, err = r.client.Collection("public").Doc(ref.ID).Set(ctx, req)
	if err != nil {
		return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firestore.SendPublic")
	}

	res.ID = ref.ID

	return &res, nil
}

func (r *repository) GetAllPublic(ctx context.Context, req *gratitude.GetAllPublicRequest) (*gratitude.GetAllPublicResponse, error) {
	var res gratitude.GetAllPublicResponse

	iter := r.client.Collection("public").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errs.Wrap(err, "repository.Firestore.GetAllPublic")
		}
		msg := gratitude.Message{}
		if err := doc.DataTo(&msg); err != nil {
			return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firestore.GetAllPublic 1")
		}
		msg.ID = doc.Ref.ID
		res.Messages = append(res.Messages, msg)
	}

	return &res, nil
}

func (r *repository) GetAllInbox(ctx context.Context, req *gratitude.GetAllInboxRequest) (*gratitude.GetAllInboxResponse, error) {
	var res gratitude.GetAllInboxResponse

	iter := r.client.Collection("users").Doc(req.UID).Collection("inbox").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firestore.GetAllInbox")
		}

		msg := gratitude.Message{}
		if err := doc.DataTo(&msg); err != nil {
			return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firestore.GetAllInbox")
		}
		msg.ID = doc.Ref.ID
		res.Messages = append(res.Messages, msg)
	}

	return &res, nil
}

func (r *repository) GetAllOutbox(ctx context.Context, req *gratitude.GetAllOutboxRequest) (*gratitude.GetAllOutboxResponse, error) {
	var res gratitude.GetAllOutboxResponse

	iter := r.client.Collection("users").Doc(req.UID).Collection("outbox").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firestore.GetAllInbox")
		}

		msg := gratitude.Message{}
		if err := doc.DataTo(&msg); err != nil {
			return nil, errs.Wrap(gratitude.NewGratitudeError(http.StatusInternalServerError, err.Error()), "repository.Firestore.GetAllInbox")
		}
		msg.ID = doc.Ref.ID
		res.Messages = append(res.Messages, msg)
	}

	return &res, nil
}
