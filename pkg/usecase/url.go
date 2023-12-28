package usecase

import (
	"context"
	"errors"
	"log"
	"short-url/pkg/domain"
	"short-url/pkg/interfaces/repository"
)

type URLInteractor struct {
	URL repository.URL
}

func (ui URLInteractor) SaveURL(ctx context.Context, url domain.URL) error {
	err := ui.URL.Save(ctx, url)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
		return errors.New("unable to save url")
	}
	return nil
}

func (ui URLInteractor) FetchURL(ctx context.Context, url domain.URL) (domain.URL, error) {
	urlBody, err := ui.URL.Fetch(ctx, url)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
		return domain.URL{}, errors.New("unable to fetch url")
	}
	return urlBody, nil
}
