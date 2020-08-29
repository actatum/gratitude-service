package transport

import (
	"os"

	"github.com/actatum/gratitude-board-service/pkg/gratitude"
	"github.com/actatum/gratitude-board-service/pkg/provider/firebase"
	"github.com/actatum/gratitude-board-service/pkg/repository/firestore"
	errs "github.com/pkg/errors"
)

// Run starts the http server
func Run() error {
	provider, err := firebase.NewFirebaseProvider()
	if err != nil {
		return errs.Wrap(err, "transport.http.Run")
	}

	repo, err := firestore.NewFirestoreRepository()
	if err != nil {
		return errs.Wrap(err, "transport.http.Run")
	}

	service := gratitude.NewGratitudeService(provider, repo)

	server := NewServer(service)

	r := routes(server)

	port := os.Getenv("PORT")

	return r.Start(":" + port)
}
