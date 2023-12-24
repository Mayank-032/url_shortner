package repository

import (
	"context"
	"short-url/pkg/domain"
)

type UrlRepo struct {
}

func (ur UrlRepo) Save(ctx context.Context, url domain.URL) error {
	return nil
}

func (ur UrlRepo) Fetch(ctx context.Context, url domain.URL) (domain.URL, error) {
	return domain.URL{}, nil
}
