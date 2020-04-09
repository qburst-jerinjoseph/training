package main

import (
	"training/internal/app"

	_ "github.com/lib/pq"
)

func main() {

	s := app.NewServer()
	s.Start()
}
