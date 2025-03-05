package main

import (
	"api/cmd/internal/config"
	"api/cmd/internal/postgresrepo"
	"api/cmd/internal/routes"
	_ "api/docs"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// @title Teste técnico
// @version 1.0
// @description Documentation for Teste técnico
// @host localhost/api
// @BasePath /
func main() {

	loadEnv()
	dsn := buildDSN()

	dbpool, err := config.OpenDB(dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer dbpool.Close()

	fmt.Println("Connected to database")
	startServer(dbpool)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func buildDSN() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func startServer(dbpool *pgxpool.Pool) {
	port := ":3000"
	fmt.Println("Starting server on port", port)

	queries := postgresrepo.New(dbpool)
	r := routes.Routes(queries)

	err := http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}
}
