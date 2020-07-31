package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// Authenticator middleware retrieves the jwt string from the 'Authorization' header.
// then calls to firebase to verify the jwt. If jwt is invalid it returns a http status 401 (unauthorized)
// If the jwt is valid it pushes the firebase auth token object into the context at the key 'token' and pushes
// the raw jwt string into the context at the key 'rawJWT'
func Authenticator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader, err := fromHeader(ctx)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}

		jwt, valid, err := verifyToken(authHeader)
		if !valid {
			log.Println(err)
			ctx.JSON(http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}

		ctx.Set("token", jwt)
		ctx.Set("rawJWT", authHeader)
		ctx.Next()
	}
}

// fromHeader retreives the 'Authorization' header from the gin context request header
func fromHeader(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return "", errors.New("missing or empty 'Authorization' header")
	}

	return token, nil
}

// getClient creates and returns a firebase auth client
func getClient() (*auth.Client, error) {
	ctx := context.Background()
	pid := os.Getenv("PROJECT_ID")
	creds := os.Getenv("GCP_KEY")
	opt := option.WithCredentialsJSON([]byte(creds))
	conf := &firebase.Config{ProjectID: pid}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// verifyToken calls to the firebase clients VerifyIDToken to verify jwt's
func verifyToken(token string) (*auth.Token, bool, error) {
	ctx := context.Background()
	client, err := getClient()
	if err != nil {
		return nil, false, err
	}

	jwt, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, false, err
	}

	return jwt, true, nil
}
