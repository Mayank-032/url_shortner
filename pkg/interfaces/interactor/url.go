package interactor

import (
	"context"
	"short-url/pkg/domain"
)

type URLInteractor interface {
	ShortURLMapper(ctx context.Context, url domain.URL) (domain.URL, error)
	FetchLongURL(ctx context.Context, url domain.URL) (domain.URL, error)
}
