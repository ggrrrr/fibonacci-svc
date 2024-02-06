package main

import (
	"fmt"
	"os"

	"github.com/ggrrrr/fibonacci-svc/common/system"
	"github.com/ggrrrr/fibonacci-svc/internal/api"
	"github.com/ggrrrr/fibonacci-svc/internal/app"
	"github.com/ggrrrr/fibonacci-svc/internal/repo/redisrepo"
)

func main() {
	s := system.NewSystem()

	repo, err := redisrepo.New(redisrepo.Config{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	app, err := app.New(repo)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	api.Register(s.Router(), app)
	s.Start()
}
