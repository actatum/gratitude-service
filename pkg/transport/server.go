package transport

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/actatum/gratitude-board-service/pkg/gratitude"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Server interface for handling http requests
type Server struct {
	service gratitude.Service
}

// NewServer returns a new server object with a gratitude service and zap logger attached to it
func NewServer(s gratitude.Service) *Server {
	return &Server{
		service: s,
	}
}

// HandleSendPrivate godoc
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
func (s *Server) HandleSendPrivate(c echo.Context) error {
	var req gratitude.Message
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	jwt, ok := c.Get("token").(*auth.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "interface is not of type *auth.Token")
	}

	req.Sender.UID = jwt.UID

	if err := validateMessage(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := s.service.SendPrivate(context.Background(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

// HandleSendPublic godoc
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
func (s *Server) HandleSendPublic(c echo.Context) error {
	var req gratitude.Message
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	jwt, ok := c.Get("token").(*auth.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "interface is not of type *auth.Token")
	}

	req.Sender.UID = jwt.UID

	if err := validateMessage(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := s.service.SendPublic(context.Background(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

// HandleGetPublic godoc
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
func (s *Server) HandleGetPublic(c echo.Context) error {
	var req gratitude.GetAllPublicRequest
	jwt, ok := c.Get("token").(*auth.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "interface is not of type *auth.Token")
	}

	req.UID = jwt.UID

	res, err := s.service.GetAllPublic(context.Background(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

// HandleGetInbox godoc
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
func (s *Server) HandleGetInbox(c echo.Context) error {
	var req gratitude.GetAllInboxRequest
	jwt, ok := c.Get("token").(*auth.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "interface is not of type *auth.Token")
	}

	req.UID = jwt.UID

	res, err := s.service.GetAllInbox(context.Background(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

// HandleGetOutbox godoc
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
func (s *Server) HandleGetOutbox(c echo.Context) error {
	var req gratitude.GetAllOutboxRequest
	jwt, ok := c.Get("token").(*auth.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "interface is not of type *auth.Token")
	}

	req.UID = jwt.UID

	res, err := s.service.GetAllOutbox(context.Background(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

// HandleGetUsers godoc
func (s *Server) HandleGetUsers(c echo.Context) error {
	var req gratitude.GetAllUsersRequest
	jwt, ok := c.Get("token").(*auth.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "interface is not of type *auth.Token")
	}

	req.UID = jwt.UID

	res, err := s.service.GetAllUsers(context.Background(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

// HandleHealth godoc
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
func (s *Server) HandleHealth(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}
