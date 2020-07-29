package api

import (
	"log"
	"net/http"
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
		return errs.Wrap(err, "api.http.Run")
	}
	repo, err := firestore.NewFirestoreRepository()
	if err != nil {
		return errs.Wrap(err, "api.http.Run")
	}
	service := gratitude.NewGratitudeService(provider, repo)
	server := NewServer(service)
	r := routes(server)

	port := os.Getenv("PORT")
	log.Println("serving application at port :" + port)
	return http.ListenAndServe(":"+port, r)
}
