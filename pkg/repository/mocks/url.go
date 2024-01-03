package mocks

import (
	"context"
	"short-url/pkg/domain"

	"github.com/stretchr/testify/mock"
)

type MockURLRepo struct {
	mock.Mock
}

func (mur *MockURLRepo) Save(ctx context.Context, url domain.URL) error {
	args := mur.Called(ctx, url)
	return args.Error(0)
}

func (mur *MockURLRepo) Fetch(ctx context.Context, url domain.URL) (domain.URL, error) {
	args := mur.Called()
	return args.Get(0).(domain.URL), args.Error(1)
}
