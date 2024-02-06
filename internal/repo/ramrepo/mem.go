package ramrepo

import (
	"sync"

	"github.com/ggrrrr/fibonacci-svc/internal/fi"
	"github.com/ggrrrr/fibonacci-svc/internal/repo"
)

type (
	memRepo struct {
		current fi.Fi
		lock    sync.Mutex
	}
)

var _ (repo.Repo) = (*memRepo)(nil)

func NewMemRepo() *memRepo {
	return &memRepo{
		current: fi.Fi{
			Previous: 0,
			Current:  0,
		},
		lock: sync.Mutex{},
	}
}

func (r *memRepo) Get() (fi.Fi, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.current, nil
}

func (*memRepo) Initialize() error {
	return nil
}

func (r *memRepo) Set(number fi.Fi) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.current = number
	return nil
}
