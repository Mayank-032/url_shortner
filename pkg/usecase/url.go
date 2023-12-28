package usecase

import (
	"context"
	"errors"
	"log"
	"short-url/pkg/domain"
	"short-url/pkg/interfaces/repository"
	"time"

	"github.com/redis/go-redis/v9"
)

type URLInteractor struct {
	URL         repository.URL
	RedisClient *redis.Client
}

func (ui URLInteractor) SaveURL(ctx context.Context, url domain.URL) error {
	key := url.Key
	if url.IsKeySigned {
		key = "-" + key
	}

	_, err := ui.RedisClient.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		if err == redis.Nil {
			return errors.New("duplicate_request")
		}

		return errors.New("unable to save url")
	}

	err = ui.URL.Save(ctx, url)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		if err.Error() == "duplicate_request" {
			return err
		}

		return errors.New("unable to save url")
	}

	ui.RedisClient.Set(ctx, key, url.LongURL, time.Hour*24)
	return nil
}

func (ui URLInteractor) FetchURL(ctx context.Context, url domain.URL) (domain.URL, error) {
	key := url.Key
	if url.IsKeySigned {
		key = "-" + key
	}

	val, err := ui.RedisClient.Get(ctx, key).Result()
	if err == nil {
		url.LongURL = val
		return url, nil
	}

	log.Println("Error: " + err.Error())

	urlBody, err := ui.URL.Fetch(ctx, url)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		if err.Error() == "invalid_hash" {
			return domain.URL{}, err
		}

		return domain.URL{}, errors.New("unable to fetch url")
	}

	ui.RedisClient.Set(ctx, key, urlBody.LongURL, time.Hour*24)
	return urlBody, nil
}
