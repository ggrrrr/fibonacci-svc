package repo

//go:generate mockgen -source=$GOFILE -destination=repo_mock.go -package=repo

import (
	"fmt"

	"github.com/ggrrrr/fibonacci-svc/internal/fi"
)

type (
	NotInitializedError string

	Config struct {
		RepoType string `envconfig:"REPO_TYPE"`
		Host     string `envconfig:"REPO_HOST"`
		Port     int    `envconfig:"REPO_PORT"`
		Username string `envconfig:"REPO_USERNAME"`
		Password string `envconfig:"REPO_PASSWORD"`
		Database string `envconfig:"REPO_DATABASE"`
		SSLMode  string `envconfig:"REPO_SSLMODE"`
	}

	Repo interface {
		Get() (fi.Fi, error)
		Set(fi fi.Fi) error
		Initialize() error
		Cleanup() error
	}
)

var EmptyNotInitializedError NotInitializedError = ""

func (e NotInitializedError) Error() string {
	return fmt.Sprintf("%s not initialized", string(e))
}
