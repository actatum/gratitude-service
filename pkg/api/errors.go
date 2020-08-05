package api

import (
	"log"

	"github.com/actatum/gratitude-board-service/pkg/gratitude"
	"github.com/gin-gonic/gin"
	errs "github.com/pkg/errors"
)

type httpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func handleHTTPError(ctx *gin.Context, err error) {
	var e *gratitude.Error
	log.Println(err)
	if errs.As(err, &e) {
		ctx.JSON(e.Status(), &httpError{
			Code:    e.Status(),
			Message: e.Error(),
		})
	}
}
