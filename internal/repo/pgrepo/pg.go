package pgrepo

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/ggrrrr/fibonacci-svc/common/log"
	"github.com/ggrrrr/fibonacci-svc/internal/fi"
	"github.com/ggrrrr/fibonacci-svc/internal/repo"
)

const RepoType string = "pg"

type (
	pgRepo struct {
		db *sql.DB
	}
)

var _ (repo.Repo) = (*pgRepo)(nil)

func New(cfg repo.Config) (*pgRepo, error) {
	if cfg.RepoType != RepoType {
		return nil, errors.New("RepoType is not pg")
	}
	if cfg.SSLMode == "" {
		cfg.SSLMode = "disable"
	}
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database, cfg.SSLMode)
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		log.Error(err).Msg("Connect")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Error(err).Msg("Ping")
		return nil, err
	}
	log.Info().
		Any("Host", cfg.Host).
		Any("Database", cfg.Database).
		Any("Port", cfg.Port).
		Msg("PostgresRepo")

	return &pgRepo{
		db: db,
	}, nil
}

func (r *pgRepo) Set(number fi.Fi) error {
	sql := `insert into fi_store( previous, current) values($1, $2)`
	_, err := r.db.Exec(sql, number.Previous, number.Current)
	if err != nil {
		return err
	}
	return nil
}

func (r *pgRepo) Get() (fi.Fi, error) {
	sql := `select  previous, current from fi_store order by id desc limit 1`
	rows, err := r.db.Query(sql)
	if err != nil {
		return fi.Fi{}, err
	}

	if rows.Err() != nil {
		return fi.Fi{}, err
	}

	var lastFi fi.Fi
	found := false

	defer rows.Close()
	for rows.Next() {
		if rows.Err() != nil {
			return fi.Fi{}, err
		}
		err = rows.Scan(&lastFi.Previous, &lastFi.Current)
		if err != nil {
			return fi.Fi{}, err
		}
		found = true
	}
	if !found {
		return fi.Fi{}, repo.NotInitializedError("pg")
	}
	return lastFi, nil
}

func (r *pgRepo) Initialize() error {
	return r.Set(fi.Fi{Previous: 0, Current: 0})
}
