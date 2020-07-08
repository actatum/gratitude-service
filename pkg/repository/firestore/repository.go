package firestore

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
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

type repository struct {
	client *firestore.Client
}

func newFirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()
	creds := os.Getenv("GCP")
	opt := option.WithCredentialsJSON([]byte(creds))
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, errs.Wrap(err, "repository.Firestore.newFirestoreClient")
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "repository.Firestore.newFirestoreClient")
	}

	return client, nil
}

// NewFirestoreRepository creates a new firestore client and then
// returns an object satisfying the gratitude.Repository interface and an error
func NewFirestoreRepository() (gratitude.Repository, error) {
	repo := &repository{}
	client, err := newFirestoreClient()
	if err != nil {
		return nil, errs.Wrap(err, "repository.Firestore.NewFirestoreRepository")
	}

	repo.client = client

	return repo, nil
}

func (r *repository) SendPrivate(ctx context.Context, req *gratitude.Message) (*gratitude.MessageResponse, error) {
	var res gratitude.MessageResponse
	jwt := ctx.Value("token").(*auth.Token)

	// Add to senders outbox
	ref, _, err := r.client.Collection("users").Doc(jwt.UID).Collection("outbox").Add(ctx, req)
	if err != nil {
		return nil, errs.Wrap(err, "repository.Firestore.SendPrivate")
	}

	// Add to all recipients inbox
	for _, recipientID := range req.RecipientIds {
		_, err = r.client.Collection("users").Doc(recipientID).Collection("inbox").Doc(ref.ID).Set(ctx, req)
		if err != nil {
			return nil, errs.Wrap(err, "repository.Firestore.SendPublic")
		}
	}

	res.MessageId = ref.ID

	return &res, nil
}

func (r *repository) SendPublic(ctx context.Context, req *gratitude.Message) (*gratitude.MessageResponse, error) {
	var res gratitude.MessageResponse
	jwt := ctx.Value("token").(*auth.Token)

	// Add to senders outbox
	ref, _, err := r.client.Collection("users").Doc(jwt.UID).Collection("outbox").Add(ctx, req)
	if err != nil {
		return nil, errs.Wrap(err, "repository.Firestore.SendPublic")
	}

	// Add to all recipients inbox
	for _, recipientID := range req.RecipientIds {
		_, err = r.client.Collection("users").Doc(recipientID).Collection("inbox").Doc(ref.ID).Set(ctx, req)
		if err != nil {
			return nil, errs.Wrap(err, "repository.Firestore.SendPublic")
		}
	}

	// Add to public message collection
	_, err = r.client.Collection("public").Doc(ref.ID).Set(ctx, req)
	if err != nil {
		return nil, errs.Wrap(err, "repository.Firestore.SendPublic")
	}

	res.MessageId = ref.ID

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
		msg := &gratitude.Message{}
		if err := doc.DataTo(msg); err != nil {
			return nil, errs.Wrap(err, "repository.Firestore.GetAllPublic")
		}
		msg.MessageId = doc.Ref.ID
		res.Messages = append(res.Messages, msg)
	}

	return &res, nil
}

func (r *repository) GetAllInbox(ctx context.Context, req *gratitude.GetAllInboxRequest) (*gratitude.GetAllInboxResponse, error) {
	var res gratitude.GetAllInboxResponse
	jwt := ctx.Value("token").(*auth.Token)

	iter := r.client.Collection("users").Doc(jwt.UID).Collection("inbox").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errs.Wrap(err, "repository.Firestore.GetAllInbox")
		}

		msg := &gratitude.Message{}
		if err := doc.DataTo(msg); err != nil {
			return nil, errs.Wrap(err, "repository.Firestore.GetAllInbox")
		}
		msg.MessageId = doc.Ref.ID
		res.Messages = append(res.Messages, msg)
	}

	return &res, nil
}

func (r *repository) GetAllOutbox(ctx context.Context, req *gratitude.GetAllOutboxRequest) (*gratitude.GetAllOutboxResponse, error) {
	var res gratitude.GetAllOutboxResponse
	jwt := ctx.Value("token").(*auth.Token)

	iter := r.client.Collection("users").Doc(jwt.UID).Collection("outbox").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errs.Wrap(err, "repository.Firestore.GetAllInbox")
		}

		msg := &gratitude.Message{}
		if err := doc.DataTo(msg); err != nil {
			return nil, errs.Wrap(err, "repository.Firestore.GetAllInbox")
		}
		msg.MessageId = doc.Ref.ID
		res.Messages = append(res.Messages, msg)
	}

	return &res, nil
}
