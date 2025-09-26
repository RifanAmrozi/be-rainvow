package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/rifanamrozi/be-rainvow/internal/config"
	"github.com/rifanamrozi/be-rainvow/internal/delivery/http"
	"github.com/rifanamrozi/be-rainvow/internal/repository"
	"github.com/rifanamrozi/be-rainvow/internal/usecase"

	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Build connection string from env vars
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	log.Println("Connected to Supabase Postgres!")

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
