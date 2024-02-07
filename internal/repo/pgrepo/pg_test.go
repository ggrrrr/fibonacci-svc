package pgrepo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ggrrrr/fibonacci-svc/internal/fi"
	"github.com/ggrrrr/fibonacci-svc/internal/repo"
)

func TestConn(t *testing.T) {
	testRepo, err := New(repo.Config{
		RepoType: RepoType,
		Host:     "localhost",
		Port:     5432,
		Username: "root",
		Password: "root",
		Database: "test",
		SSLMode:  "",
	})

	require.NoError(t, err)
	require.NotNil(t, testRepo)

	_, err = testRepo.db.Exec(`truncate table fi_store`)
	require.NoError(t, err)

	_, err = testRepo.Get()
	require.Error(t, err)
	require.True(t, errors.As(err, &repo.EmptyNotInitializedError))

	err = testRepo.Initialize()
	require.NoError(t, err)

	v0, err := testRepo.Get()
	require.NoError(t, err)
	assert.Equal(t, fi.Fi{Previous: 0, Current: 0}, v0)

	fi1 := fi.Fi{
		Previous: 0,
		Current:  0,
	}

	err = testRepo.Set(fi1)
	require.NoError(t, err)

	v1, err := testRepo.Get()
	require.NoError(t, err)
	require.Equal(t, fi1, v1)

}
