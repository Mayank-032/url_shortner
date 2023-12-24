package repository

import (
	"context"
	"short-url/pkg/domain"
)

type URL interface {
	Save(ctx context.Context, url domain.URL) error
	Fetch(ctx context.Context, url domain.URL) (domain.URL, error)
}
