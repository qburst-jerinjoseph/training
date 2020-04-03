package app

import (
	"log"
	"training/internal/data"

	"github.com/sirupsen/logrus"
)

type server struct {
	data.Repo
	*logrus.Logger
}

func NewServer() *server {
	logger := logrus.New()
	repo, err := data.NewPostgresRepo()
	if err != nil {
		log.Fatal("error while connecting to database", err)
	}
	return &server{
		repo,
		logger,
	}
}
