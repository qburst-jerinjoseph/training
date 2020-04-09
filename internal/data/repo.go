package data

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Repo is
type Repo interface {
	AlterationMethodList() ([]AlterationMethod, error)
	AlterationMethodGet(id int) (*AlterationMethod, error)
}
type postgresRepo struct {
	*sql.DB
}

func NewPostgresRepo() (Repo, error) {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "training_develop"
	}
	connectionStr := "postgres://training:training@localhost:5432/%s?sslmode=disable"

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			connectionStr,
			dbName,
		),
	)
	if err != nil {
		return nil, err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, errors.Wrap(
			err,
			"can not create driver instance",
		)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///home/jerin/Documents/projects/go/src/training/internal/data/migration",
		"postgres", driver)
	if err != nil {
		return nil, errors.Wrap(
			err,
			"can not create database instance",
		)
	}
	m.Up()

	return &postgresRepo{db}, nil
}
