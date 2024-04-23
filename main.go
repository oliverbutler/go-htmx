package main

import (
	"api/app_api"
	"api/lib"
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load .env file: %v, using environment variables", err)
	}

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}

	app := lib.NewApp(pool)

	defer pool.Close()

	r := gin.Default()

	r.Static("/assets", "./assets")

	r.ForwardedByClientIP = true
	err = r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}

	app_api.InitAppAPI(r, app)

	host := os.Getenv("HOST")

	if host == "" {
		host = "localhost:8069"
	}

	err = r.Run(host)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
