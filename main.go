package main

import (
	"net/http"
	"os"
	"training/internal/app"

	_ "github.com/lib/pq"
)

func main() {

	s := app.NewServer()
	s.Info("starting server on port: 3005")
	s.Info("Database: ", os.Getenv("POSTGRES_DB"))
	err := http.ListenAndServe(":3005", s.InitRouter())
	if err != nil {
		s.Fatal("Error while starting server", err)
	}
}
