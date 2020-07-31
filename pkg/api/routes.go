package api

import (
	"github.com/actatum/gratitude-board-service/pkg/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func routes(s Server) *gin.Engine {
	r := gin.Default()
	r.Use(corsOptions())
	r.GET("/api/v1/health", s.HandleHealth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	v1.Use(middleware.Authenticator())
	{
		public := v1.Group("/public")
		{
			public.GET("", s.HandleGetPublic)
			public.POST("", s.HandleSendPublic)
		}

		private := v1.Group("/private")
		{
			private.POST("", s.HandleSendPrivate)
		}

		inbox := v1.Group("/inbox")
		{
			inbox.GET("", s.HandleGetInbox)
		}

		outbox := v1.Group("/outbox")
		{
			outbox.GET("", s.HandleGetOutbox)
		}

		users := v1.Group("/users")
		{
			users.GET("", s.HandleGetUsers)
		}
	}

	return r
}

func corsOptions() gin.HandlerFunc {
	opts := cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000", "*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"ACCEPT", "Authorization", "Content-Type", "Origin"},
		ExposeHeaders: []string{"Content-Length"},
	})

	return opts
}
