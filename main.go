package main

import (
	"context"
	"fmt"

	"github.com/kelseyhightower/envconfig"

	"github.com/ggrrrr/fibonacci-svc/common/system"
	"github.com/ggrrrr/fibonacci-svc/internal/api"
	"github.com/ggrrrr/fibonacci-svc/internal/app"
	"github.com/ggrrrr/fibonacci-svc/internal/repo"
	"github.com/ggrrrr/fibonacci-svc/internal/repo/pgrepo"
	"github.com/ggrrrr/fibonacci-svc/internal/repo/ramrepo"
	"github.com/ggrrrr/fibonacci-svc/internal/repo/redisrepo"
)

type (
	SvcConfig struct {
		System system.Config
		Repo   repo.Config
	}
)

func main() {
	var cfg SvcConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		fmt.Printf("config process: %v\n", err)
		panic(1)
	}

	if cfg.System.Addr == "" {
		fmt.Printf("config system: LISTEN_ADDR empty %v\n", err)
		panic(1)
	}

	s := system.NewSystem(context.Background(), cfg.System)

	var repo repo.Repo

	switch cfg.Repo.RepoType {
	case redisrepo.RepoType:
		repo, err = redisrepo.New(cfg.Repo)
		if err != nil {
			fmt.Printf("redis error: %v\n", err)
			panic(1)
		}
	case pgrepo.RepoType:
		repo, err = pgrepo.New(cfg.Repo)
		if err != nil {
			fmt.Printf("pg error: %v\n", err)
			panic(1)
		}
	default:
		repo = ramrepo.New()
	}
	s.AddCleanup(repo.Cleanup)
	app, err := app.New(repo)
	if err != nil {
		fmt.Printf("app error: %v\n", err)
		panic(1)
	}

	api.Register(s.Router("/v1"), app)
	s.Start()
}
