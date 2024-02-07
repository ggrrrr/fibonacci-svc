package app_test

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ggrrrr/fibonacci-svc/internal/app"
	"github.com/ggrrrr/fibonacci-svc/internal/fi"
	"github.com/ggrrrr/fibonacci-svc/internal/repo"
	"github.com/ggrrrr/fibonacci-svc/internal/repo/ramrepo"
	"github.com/ggrrrr/fibonacci-svc/internal/repo/redisrepo"
)

func TestNew(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testRepo := repo.NewMockRepo(mockCtrl)

	type testCase struct {
		name      string
		prep      func(t *testing.T)
		expError  error
		returnVal bool
	}

	tests := []testCase{
		{
			name: "ok",
			prep: func(t *testing.T) {
				testRepo.EXPECT().Get().Times(1).Return(fi.Fi{}, nil)

			},
			expError:  nil,
			returnVal: true,
		},
		{
			name: "ok with init",
			prep: func(t *testing.T) {
				testRepo.EXPECT().Get().Times(1).Return(fi.Fi{}, repo.EmptyNotInitializedError)
				testRepo.EXPECT().Initialize().Times(1).Return(nil)

			},
			expError:  nil,
			returnVal: true,
		},
		{
			name: "error on get",
			prep: func(t *testing.T) {
				testRepo.EXPECT().Get().Times(1).Return(fi.Fi{}, errors.New("GetError"))
				// testRepo.EXPECT().Initialize().Times(1).Return(errors.New("SomeError"))

			},
			expError:  errors.New("GetError"),
			returnVal: false,
		},
		{
			name: "error with init",
			prep: func(t *testing.T) {
				testRepo.EXPECT().Get().Times(1).Return(fi.Fi{}, repo.EmptyNotInitializedError)
				testRepo.EXPECT().Initialize().Times(1).Return(errors.New("SomeError"))

			},
			expError:  errors.New("SomeError"),
			returnVal: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.prep(t)
			a, err := app.New(testRepo)
			if tc.returnVal {
				assert.NotNil(t, a)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, err, tc.expError)
			}
		})
	}
}

func TestApp(t *testing.T) {
	memRepo := ramrepo.New()

	testApp, err := app.New(memRepo)
	require.NoError(t, err)
	require.NotNil(t, testApp)

	wg := sync.WaitGroup{}
	wg.Add(6)
	i := 0
	for i <= 5 {
		i++
		go func() {
			defer wg.Done()
			_, err := testApp.Next()
			require.NoError(t, err)
		}()

	}
	wg.Wait()
	assert.Equal(t, fi.Number(8), testApp.Current())

	wg.Add(2)
	i = 0
	for i <= 1 {
		i++
		go func() {
			defer wg.Done()
			_, err := testApp.Previous()
			require.NoError(t, err)
		}()

	}
	wg.Wait()

	assert.Equal(t, fi.Number(3), testApp.Current())
}

func setupMem() (*app.App, error) {
	return app.New(ramrepo.New())
}

func setupRedis() (*app.App, error) {
	testRepo, err := redisrepo.New(repo.Config{
		RepoType: redisrepo.RepoType,
		Host:     "localhost",
		Port:     6379,
		Password: "",
		Database: "0",
	})
	testRepo.Initialize()
	if err != nil {
		return nil, err
	}
	return app.New(testRepo)
}

func TestMem(t *testing.T) {
	testApp, err := setupMem()
	require.NoError(t, err)
	loop := atomic.Bool{}
	loop.Store(true)
	wg := sync.WaitGroup{}
	wg.Add(1)
	ticker := time.NewTicker(1 * time.Second)
	counter := 0
	go func() {
		<-ticker.C
		loop.Store(false)
	}()
	func() {
		defer wg.Done()
		for loop.Load() {
			testApp.Next()
			counter++
		}
	}()
	wg.Wait()
	fmt.Printf("ASDASD %v %v\n", counter, testApp.Current())
}

func TestRedis(t *testing.T) {
	testApp, err := setupRedis()
	require.NoError(t, err)
	loop := atomic.Bool{}
	loop.Store(true)
	wg := sync.WaitGroup{}
	wg.Add(1)
	ticker := time.NewTicker(1 * time.Second)
	counter := 0
	go func() {
		<-ticker.C
		loop.Store(false)
	}()
	func() {
		defer wg.Done()
		for loop.Load() {
			testApp.Next()
			counter++
		}
	}()
	wg.Wait()
	fmt.Printf("ASDASD %v %v\n", counter, testApp.Current())
}

func BenchmarkRedis(b *testing.B) {
	testApp, err := setupRedis()
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testApp.Next()
	}
}
