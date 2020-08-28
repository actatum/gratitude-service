package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CorsDefault returns default cors config for all services
func CorsDefault() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType},
	}
}
