package redisrepo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ggrrrr/fibonacci-svc/internal/fi"
	"github.com/ggrrrr/fibonacci-svc/internal/repo"
)

func TestSetGet(t *testing.T) {
	var err error

	testRepo, err := New(Config{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	require.NoError(t, err)
	require.NotNil(t, testRepo)

	// Remove key
	_ = testRepo.client.Del(testRepo.redisKey)

	// Verify empty key error
	_, err = testRepo.Get()
	require.Error(t, err)
	require.True(t, errors.As(err, &repo.EmptyNotInitializedError))

	err = testRepo.Initialize()
	require.NoError(t, err)

	v0, err := testRepo.Get()
	require.NoError(t, err)
	assert.Equal(t, fi.Fi{Previous: 0, Current: 0}, v0)

	fi1 := fi.Fi{
		Previous: 3,
		Current:  5,
	}

	err = testRepo.Set(fi1)
	require.NoError(t, err)

	v1, err := testRepo.Get()
	require.NoError(t, err)

	assert.Equal(t, fi1, v1)
}
