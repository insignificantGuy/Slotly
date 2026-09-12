package apiserver

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	api "github.com/insignificantGuy/Slotly/internal/api"
	"github.com/insignificantGuy/Slotly/internal/auth"
	"github.com/insignificantGuy/Slotly/internal/database"
)

var (
	svc        *api.Service
	controller *api.Controller
	tokens     *auth.Tokens
)

func Start() error {
	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	router, err := createRouter()
	if err != nil {
		return err
	}

	fmt.Printf("Slotly API listening on %s\n", addr)
	return router.Run(addr)
}

func createRouter() (*gin.Engine, error) {
	db, err := database.InitMySQL()
	if err != nil {
		return nil, err
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "dev-insecure-jwt-secret"
	}
	tokens = auth.NewTokens(secret, 15*time.Minute)

	svc, err = api.NewService(db, tokens)
	if err != nil {
		return nil, err
	}

	controller = api.NewController(svc)

	router := gin.Default()
	SetupRoutes(router, tokens)
	return router, nil
}
