package api

import (
	"github.com/gin-gonic/gin"
)

// Server interface for handling http requests
type Server interface {
	HandleSendPrivate(*gin.Context)
	HandleSendPublic(*gin.Context)
	HandleGetPublic(*gin.Context)
	HandleGetInbox(*gin.Context)
	HandleGetOutbox(*gin.Context)
	HandleGetUsers(*gin.Context)
	HandleHealth(*gin.Context)
}
