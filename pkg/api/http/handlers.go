package http

import (
	"encoding/json"
	"github.com/go-chi/render"
	"log"
	"net/http"

	gratitude "github.com/actatum/gratitude-board-service/pkg/service"
)

// Server interface for handling http requests
type Server interface {
	HandleSendPrivate() http.HandlerFunc
	HandleSendPublic() http.HandlerFunc
	HandleGetPublic() http.HandlerFunc
	HandleGetInbox() http.HandlerFunc
	HandleGetOutbox() http.HandlerFunc
	HandleGetUsers() http.HandlerFunc
}

type server struct {
	service gratitude.GratitudeServiceServer
}

// NewServer returns an implementation of the Server interface for handling http requests
func NewServer(service gratitude.GratitudeServiceServer) Server {
	return &server{
		service: service,
	}
}

func (s *server) HandleSendPrivate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req gratitude.Message
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			handleHttpError(w, err, http.StatusBadRequest)
			return
		}

		res, err := s.service.SendPrivate(r.Context(), &req)
		if err != nil {
			handleHttpError(w, err, http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, res)
	}
}

func (s *server) HandleSendPublic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req gratitude.Message
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			handleHttpError(w, err, http.StatusBadRequest)
			return
		}

		res, err := s.service.SendPublic(r.Context(), &req)
		if err != nil {
			handleHttpError(w, err, http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, res)
	}
}

func (s *server) HandleGetPublic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := s.service.GetAllPublic(r.Context(), nil)
		if err != nil {
			handleHttpError(w, err, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, res)
	}
}

func (s *server) HandleGetInbox() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := s.service.GetAllInbox(r.Context(), nil)
		if err != nil {
			handleHttpError(w, err, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, res)
	}
}

func (s *server) HandleGetOutbox() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := s.service.GetAllOutbox(r.Context(), nil)
		if err != nil {
			handleHttpError(w, err, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, res)
	}
}

func (s *server) HandleGetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := s.service.GetAllUsers(r.Context(), nil)
		if err != nil {
			handleHttpError(w, err, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, res)
	}
}

func handleHttpError(w http.ResponseWriter, err error, status int) {
	log.Println(err)
	http.Error(w, err.Error(), status)
}
