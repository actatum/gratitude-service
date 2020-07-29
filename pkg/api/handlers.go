package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/actatum/gratitude-board-service/pkg/gratitude"
)

type server struct {
	service gratitude.GratitudeService
}

// NewServer returns an implementation of the Server interface for handling http requests
func NewServer(service gratitude.GratitudeService) Server {
	return &server{
		service: service,
	}
}

func (s *server) HandleSendPrivate(ctx *gin.Context) {
	var req gratitude.Message
	if err := ctx.BindJSON(&req); err != nil {
		handleHttpError(ctx, gratitude.NewGratitudeError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := validateMessage(&req); err != nil {
		handleHttpError(ctx, err)
		return
	}

	res, err := s.service.SendPrivate(ctx, &req)
	if err != nil {
		handleHttpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (s *server) HandleSendPublic(ctx *gin.Context) {
	var req gratitude.Message
	if err := ctx.BindJSON(&req); err != nil {
		handleHttpError(ctx, gratitude.NewGratitudeError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := validateMessage(&req); err != nil {
		handleHttpError(ctx, err)
		return
	}

	res, err := s.service.SendPublic(ctx, &req)
	if err != nil {
		handleHttpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (s *server) HandleGetPublic(ctx *gin.Context) {
	res, err := s.service.GetAllPublic(ctx, &gratitude.GetAllPublicRequest{})
	if err != nil {
		handleHttpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (s *server) HandleGetInbox(ctx *gin.Context) {
	res, err := s.service.GetAllInbox(ctx, &gratitude.GetAllInboxRequest{})
	if err != nil {
		handleHttpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (s *server) HandleGetOutbox(ctx *gin.Context) {
	res, err := s.service.GetAllOutbox(ctx, &gratitude.GetAllOutboxRequest{})
	if err != nil {
		handleHttpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (s *server) HandleGetUsers(ctx *gin.Context) {
	res, err := s.service.GetAllUsers(ctx, &gratitude.GetAllUsersRequest{})
	if err != nil {
		handleHttpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (s *server) HandleHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "healthy")
}
