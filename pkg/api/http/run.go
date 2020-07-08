package http

import (
	"github.com/actatum/gratitude-board-service/pkg/provider/firebase"
	"github.com/actatum/gratitude-board-service/pkg/repository/firestore"
	gratitude "github.com/actatum/gratitude-board-service/pkg/service"
	errs "github.com/pkg/errors"
	"log"
	"net/http"
	"os"
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
	if err := walk(r); err != nil {
		return errs.Wrap(err, "api.http.Run")
	}

	port := os.Getenv("PORT")
	log.Println("serving application at port :" + port)
	return http.ListenAndServe(":"+port, r)
}
