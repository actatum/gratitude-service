package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/actatum/gratitude-board-service/pkg/gratitude"
)

type server struct {
	service gratitude.Service
}

// NewServer returns an implementation of the Server interface for handling http requests
func NewServer(service gratitude.Service) Server {
	return &server{
		service: service,
	}
}

// SendPrivate godoc
// @Summary Send a user a private message
// @Description send a user a private message
// @ID Send Private
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Param message body gratitude.Message true "message"
// @Success 200 {object} gratitude.SendResponse
// @Failure 400 {object} httpError
// @Failure 500 {object} httpError
// @Router /private [post]
func (s *server) HandleSendPrivate(ctx *gin.Context) {
	var req gratitude.Message
	if err := ctx.BindJSON(&req); err != nil {
		handleHTTPError(ctx, gratitude.NewGratitudeError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := validateMessage(&req); err != nil {
		handleHTTPError(ctx, err)
		return
	}

	res, err := s.service.SendPrivate(ctx, &req)
	if err != nil {
		handleHTTPError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// SendPublic godoc
// @Summary Send a public message
// @Description send a public message
// @ID Send Private
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Param message body gratitude.Message true "message"
// @Success 200 {object} gratitude.SendResponse
// @Failure 400 {object} httpError
// @Failure 500 {object} httpError
// @Router /public [post]
func (s *server) HandleSendPublic(ctx *gin.Context) {
	var req gratitude.Message
	if err := ctx.BindJSON(&req); err != nil {
		handleHTTPError(ctx, gratitude.NewGratitudeError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := validateMessage(&req); err != nil {
		handleHTTPError(ctx, err)
		return
	}

	res, err := s.service.SendPublic(ctx, &req)
	if err != nil {
		handleHTTPError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetPublic godoc
// @Summary Retrieve all public messages
// @Description retrieve all public messages
// @ID Get Public
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} gratitude.GetAllPublicResponse
// @Failure 400 {object} httpError
// @Failure 500 {object} httpError
// @Router /public [get]
func (s *server) HandleGetPublic(ctx *gin.Context) {
	res, err := s.service.GetAllPublic(ctx, &gratitude.GetAllPublicRequest{})
	if err != nil {
		handleHTTPError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetInbox godoc
// @Summary Retrieve a users inbox
// @Description retrieve all messages in a users inbox
// @ID Get Inbox
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} gratitude.GetAllInboxResponse
// @Failure 400 {object} httpError
// @Failure 500 {object} httpError
// @Router /inbox [get]
func (s *server) HandleGetInbox(ctx *gin.Context) {
	res, err := s.service.GetAllInbox(ctx, &gratitude.GetAllInboxRequest{})
	if err != nil {
		handleHTTPError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetOutbox godoc
// @Summary Retrieve a users outbox
// @Description retrieve all messages in users outbox
// @ID Get Inbox
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} gratitude.GetAllOutboxResponse
// @Failure 400 {object} httpError
// @Failure 500 {object} httpError
// @Router /outbox [get]
func (s *server) HandleGetOutbox(ctx *gin.Context) {
	res, err := s.service.GetAllOutbox(ctx, &gratitude.GetAllOutboxRequest{})
	if err != nil {
		handleHTTPError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (s *server) HandleGetUsers(ctx *gin.Context) {
	res, err := s.service.GetAllUsers(ctx, &gratitude.GetAllUsersRequest{})
	if err != nil {
		handleHTTPError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// Health godoc
// @Summary Server health check
// @Description check the servers health status
// @ID Health
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Success 200 {string} healthy
// @Failure 400 {object} httpError
// @Failure 500 {object} httpError
// @Router /health [get]
func (s *server) HandleHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "healthy")
}
