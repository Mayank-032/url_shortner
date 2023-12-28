package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"short-url/pkg/domain"

	"github.com/go-sql-driver/mysql"
)

type UrlRepo struct {
	DB *sql.DB
}

func (ur UrlRepo) Save(ctx context.Context, url domain.URL) error {
	var mysqlErr *mysql.MySQLError

	db := ur.DB

	sqlQuery := `
		INSERT INTO url_mapper(short_url, long_url, is_hash_signed, hash) VALUES(?, ?, ?, ?)
	`

	_, err := db.ExecContext(ctx, sqlQuery, url.ShortURL, url.LongURL, url.IsKeySigned, url.Key)
	if err != nil {
		log.Println("Error: " + err.Error())

		if errors.As(err, &mysqlErr) && mysqlErr.Number == domain.ErrorDuplicateRecord {
			return errors.New("duplicate_request")
		}

		return errors.New("unable to execute sql query")
	}

	return nil
}

func (ur UrlRepo) Fetch(ctx context.Context, url domain.URL) (domain.URL, error) {
	db := ur.DB

	sqlQuery := `
		SELECT 
			long_url
		FROM url_mapper
		WHERE
			hash = ? and is_hash_signed = ?
	`

	var longURL string
	err := db.QueryRowContext(ctx, sqlQuery, url.Key, url.IsKeySigned).Scan(&longURL)
	if err != nil {
		log.Println("Error: " + err.Error())

		if err == sql.ErrNoRows {
			return domain.URL{}, errors.New("invalid_hash")
		}

		return domain.URL{}, errors.New("unable to execute query")
	}

	url.LongURL = longURL

	return url, nil
}
