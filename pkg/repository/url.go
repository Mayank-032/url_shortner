package repository

import (
	"context"
	"encoding/json"
	"log"
	"short-url/pkg/domain"
)

type UrlRepo struct {
}

func (ur UrlRepo) Save(ctx context.Context, url domain.URL) error {
	bytes, _ := json.Marshal(url)
	log.Println("request Bytes: " + string(bytes))
	
	return nil
}

func (ur UrlRepo) Fetch(ctx context.Context, url domain.URL) (domain.URL, error) {
	return domain.URL{}, nil
}
