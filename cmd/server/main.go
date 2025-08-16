// cmd/server/main.go
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/zkurdi45/mystack-go/internal/data"
	"github.com/zkurdi45/mystack-go/internal/handler"
	"github.com/zkurdi45/mystack-go/internal/router"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	db_dsn := os.Getenv("DATABASE_URL")
	if db_dsn == "" {
		logger.Fatal("FATAL: DATABASE_URL environment variable is not set")
	}
	db, err := data.OpenDB(db_dsn)
	if err != nil {
		logger.Fatalf("FATAL: could not connect to database: %s", err)
	}
	defer db.Close()
	logger.Println("database connection pool established")

	models := data.NewModels(db)

	// Increase timeout to 5 minutes to allow for long-running PDF processing.
	httpClient := &http.Client{
		Timeout: time.Minute * 5,
	}

	h := handler.New(
		logger,
		models,
	)

	r := router.New(h)
	port := getPort()
	logger.Printf("server starting on http://localhost%s", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		logger.Fatalf("could not start server: %s\n", err)
	}
}

func getPort() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
