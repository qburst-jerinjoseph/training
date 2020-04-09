package app

import (
	"log"
	"net/http"
	"os"
	"training/internal/data"

	"github.com/sirupsen/logrus"
)

type server struct {
	port string
	data.Repo
	*logrus.Logger
}

func NewServer() *server {
	logger := logrus.New()
	repo, err := data.NewPostgresRepo()
	if err != nil {
		log.Fatal("error while connecting to database", err)
	}
	s := &server{
		"",
		repo,
		logger,
	}
	s.port = os.Getenv("API_PORT")
	if len(s.port) == 0 {
		s.port = "3005"
	}
	return s
}
func (s server) Start() {
	s.Info("starting server on port:", s.port)
	err := http.ListenAndServe(":"+s.port, s.InitRouter())
	if err != nil {
		s.Fatal("Error while starting server", err)
	}
}
