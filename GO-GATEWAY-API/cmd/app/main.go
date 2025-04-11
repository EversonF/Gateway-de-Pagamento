package main

import (
	"log"
	"fmt"
	"os"
	"github.com/devfullcycle/go-gateway-api/internal/web/server"
	"github.com/devfullcycle/go-gateway-api/internal/service"
	"github.com/devfullcycle/go-gateway-api/internal/repository"
	"github.com/joho/godotenv"
	_"github.com/lib/pq"
	"database/sql"
)

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main ()  {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_SSLMODE", "disable"),

	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: ", err)
	}
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	port := getEnv("PORT", "8080")
	server := server.NewServer(accountService, port)
	server.ConfigureRoutes()

	server.Start(); if err != nil {
		log.Fatalf("Error starting server: ", err)
	}
}