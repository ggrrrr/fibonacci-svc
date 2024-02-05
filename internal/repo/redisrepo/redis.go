package redisrepo

import (
	"encoding/json"
	"errors"

	"github.com/go-redis/redis"

	"github.com/ggrrrr/fibonacci-svc/internal/fi"
	"github.com/ggrrrr/fibonacci-svc/internal/repo"
)

type (
	Config struct {
		Addr     string
		Password string
		DB       int
	}

	redisRepo struct {
		redisKey string
		client   *redis.Client
	}
)

var _ (repo.Repo) = (*redisRepo)(nil)

func New(cfg Config) (*redisRepo, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &redisRepo{
		redisKey: "last.fi",
		client:   client,
	}, nil
}

func (r *redisRepo) Set(fi fi.Fi) error {
	j, err := json.Marshal(fi)
	if err != nil {
		return err
	}
	res := r.client.Set(r.redisKey, j, 0)
	if res.Err() != nil {
		return err
	}
	return nil
}

func (r *redisRepo) Get() (fi.Fi, error) {
	val, err := r.client.Get(r.redisKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return fi.Fi{}, repo.NotInitializedError("redis")
		}
		return fi.Fi{}, err
	}
	var out fi.Fi
	err = json.Unmarshal([]byte(val), &out)
	if err != nil {
		return fi.Fi{}, err
	}
	return out, err
}

func (r *redisRepo) Initialize() error {
	return r.Set(fi.Fi{Previous: 0, Current: 0})
}
