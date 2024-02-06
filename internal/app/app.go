package app

import (
	"errors"
	"sync"

	"github.com/ggrrrr/fibonacci-svc/internal/fi"
	"github.com/ggrrrr/fibonacci-svc/internal/repo"
)

type (
	App struct {
		current fi.Fi
		mutex   sync.RWMutex
		repo    repo.Repo
	}
)

// Create App instance and initializes the repository if needed
func New(r repo.Repo) (*App, error) {
	current, err := r.Get()
	if err != nil {
		if !errors.As(err, &repo.EmptyNotInitializedError) {
			return nil, err
		}
		err = r.Initialize()
		if err != nil {
			return nil, err
		}
	}
	return &App{
		current: current,
		mutex:   sync.RWMutex{},
		repo:    r,
	}, nil
}

// Fetches/Calculates/Stores next Fi value
func (a *App) Next() (fi.Number, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	next := fi.Next(a.current)
	err := a.repo.Set(next)
	if err != nil {
		return fi.Number(0), err
	}
	a.current = next
	return next.Current, nil
}

// Fetches/Calculates/Stores Previous Fi value
func (a *App) Previous() (fi.Number, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	out := fi.Previous(a.current)
	err := a.repo.Set(out)
	if err != nil {
		return fi.Number(0), err
	}
	a.current = out
	return out.Current, nil
}

// Fetches the current/last Fi value
func (a *App) Current() fi.Number {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.current.Current
}
