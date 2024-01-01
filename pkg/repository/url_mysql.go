package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"short-url/pkg/domain"

	"github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

type UrlRepo struct {
	DB          *sql.DB
	RedisClient *redis.Client
}

func (ur UrlRepo) Save(ctx context.Context, url domain.URL) error {
	key := url.Key
	if url.IsKeySigned {
		key = "-" + key
	}

	// check if key is present in cache or not
	val, err := ur.Get(ctx, key)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
		return errors.New("unable to fetch url")
	}

	if len(val) > 0 {
		return errors.New("duplicate_request")
	}

	// if saved return duplicate error and if not save it into db, if key present will return duplicate error
	var mysqlErr *mysql.MySQLError
	db := ur.DB

	sqlQuery := `
		INSERT INTO url_mapper(short_url, long_url, is_hash_signed, hash) VALUES(?, ?, ?, ?)
	`

	_, err = db.ExecContext(ctx, sqlQuery, url.ShortURL, url.LongURL, url.IsKeySigned, url.Key)
	if err != nil {
		log.Println("Error: " + err.Error())

		if errors.As(err, &mysqlErr) && mysqlErr.Number == domain.ErrorDuplicateRecord {
			return errors.New("duplicate_request")
		}

		return errors.New("unable to execute sql query")
	}

	// if there is no error and key is new set it in our cache
	ur.Set(ctx, key, url.LongURL)
	return nil
}

func (ur UrlRepo) Fetch(ctx context.Context, url domain.URL) (domain.URL, error) {
	// First check in cache if data is present
	key := url.Key
	if url.IsKeySigned {
		key = "-" + key
	}

	val, err := ur.Get(ctx, key)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
		return domain.URL{}, errors.New("unable to fetch url")
	}
	if len(val) > 0 {
		log.Println("Data fetched from cache")
		url.LongURL = val
		return url, nil
	}

	// now fetch data from db if not present in cache
	db := ur.DB

	sqlQuery := `
		SELECT 
			long_url
		FROM url_mapper
		WHERE
			hash = ? and is_hash_signed = ?
	`

	var longURL string
	err = db.QueryRowContext(ctx, sqlQuery, url.Key, url.IsKeySigned).Scan(&longURL)
	if err != nil {
		log.Println("Error: " + err.Error())

		if err == sql.ErrNoRows {
			return domain.URL{}, errors.New("invalid_hash")
		}

		return domain.URL{}, errors.New("unable to execute query")
	}
	url.LongURL = longURL

	// now set the data in cache
	ur.Set(ctx, key, longURL)
	return url, nil
}
