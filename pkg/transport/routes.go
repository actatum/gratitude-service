package transport

import (
	mid "github.com/actatum/gratitude-board-service/pkg/transport/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
)

func routes(s *Server) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(mid.CorsDefault()))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/api/health", s.HandleHealth)

	api := e.Group("/api")
	api.Use(middleware.Logger())
	api.Use(mid.Authenticator())
	{
		public := api.Group("/public")
		{
			public.GET("", s.HandleGetPublic)
			public.POST("", s.HandleSendPublic)
		}

		private := api.Group("/private")
		{
			private.POST("", s.HandleSendPrivate)
		}

		inbox := api.Group("/inbox")
		{
			inbox.GET("", s.HandleGetInbox)
		}

		outbox := api.Group("/outbox")
		{
			outbox.GET("", s.HandleGetOutbox)
		}

		users := api.Group("/users")
		{
			users.GET("", s.HandleGetUsers)
		}
	}

	return e
}
