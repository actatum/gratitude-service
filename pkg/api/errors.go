package api

import (
	"log"

	"github.com/actatum/gratitude-board-service/pkg/gratitude"
	"github.com/gin-gonic/gin"
	errs "github.com/pkg/errors"
)

func handleHttpError(ctx *gin.Context, err error) {
	var e *gratitude.GratitudeError
	log.Println(err)
	if errs.As(err, &e) {
		ctx.JSON(e.Status(), e)
	}
}
