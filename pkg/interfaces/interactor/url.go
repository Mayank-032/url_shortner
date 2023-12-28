package interactor

import (
	"context"
	"short-url/pkg/domain"
)

type URLInteractor interface {
	SaveURL(ctx context.Context, url domain.URL) error
	FetchURL(ctx context.Context, url domain.URL) (domain.URL, error)
}
