package repo

//go:generate mockgen -source=$GOFILE -destination=repo_mock.go -package=repo

import (
	"fmt"

	"github.com/ggrrrr/fibonacci-svc/internal/fi"
)

type (
	NotInitializedError string

	Repo interface {
		Get() (fi.Fi, error)
		Set(fi fi.Fi) error
		Initialize() error
	}
)

var EmptyNotInitializedError NotInitializedError = ""

func (e NotInitializedError) Error() string {
	return fmt.Sprintf("%s not initialized", string(e))
}
