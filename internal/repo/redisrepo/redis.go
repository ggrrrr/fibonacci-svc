package redisrepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"

	"github.com/ggrrrr/fibonacci-svc/common/log"
	"github.com/ggrrrr/fibonacci-svc/internal/fi"
	"github.com/ggrrrr/fibonacci-svc/internal/repo"
)

const RepoType string = "redis"
const redisKey string = "last.fi"

type (
	redisRepo struct {
		redisKey string
		client   *redis.Client
	}
)

var _ (repo.Repo) = (*redisRepo)(nil)

func New(cfg repo.Config) (*redisRepo, error) {
	if cfg.RepoType != RepoType {
		return nil, errors.New("RepoType is not redis")
	}

	// If Database is not set we use default == 0
	db, _ := strconv.Atoi(cfg.Database)
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Error(err).Str("addr", addr).Int("db", db).Str("redisKey", redisKey).Msg("RedisRepo")
		return nil, err
	}
	log.Info().Str("addr", addr).Int("db", db).Str("redisKey", redisKey).Msg("RedisRepo")
	return &redisRepo{
		redisKey: redisKey,
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
