package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/rifanamrozi/be-rainvow/internal/config"
	"github.com/rifanamrozi/be-rainvow/internal/delivery/http"
	"github.com/rifanamrozi/be-rainvow/internal/repository"
	"github.com/rifanamrozi/be-rainvow/internal/usecase"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	// Setup repository (in-memory for now, can be swapped with DB)
	userRepo := repository.NewInMemoryUserRepository()

	// Setup usecase
	userUC := usecase.NewUserUsecase(userRepo)

	// Setup Gin
	router := gin.Default()

	// Register HTTP handlers
	http.NewUserHandler(router, userUC)
	http.NewEdgeHandler(router)

	addr := fmt.Sprintf(":%d", cfg.AppPort)
	log.Printf("server running on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
