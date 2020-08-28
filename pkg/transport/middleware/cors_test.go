package middleware

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestCorsDefault(t *testing.T) {
	tests := []struct {
		name     string
		expected middleware.CORSConfig
	}{
		{name: "default conifg", expected: middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:3000", "*"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders: []string{echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType},
		}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CorsDefault()

			assert.Equal(t, got, tc.expected)
		})
	}
}
