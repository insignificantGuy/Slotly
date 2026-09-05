package apiserver

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	api "github.com/insignificantGuy/Slotly/internal/api"
	"github.com/insignificantGuy/Slotly/internal/database"
)

var (
	svc        *api.Service
	controller *api.Controller
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

	svc, err = api.NewService(db)
	if err != nil {
		return nil, err
	}

	controller = api.NewController(svc)

	router := gin.Default()
	SetupRoutes(router)
	return router, nil
}
