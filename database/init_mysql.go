package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"short-url/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMySQL() error {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.Configuration.Database.Username, config.Configuration.Database.Password, config.Configuration.Database.Host, config.Configuration.Database.Port, config.Configuration.Database.Schema)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("Error: " + err.Error())
		return errors.New("unable to open db conn")
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	err = db.PingContext(ctx)
	if err != nil {
		log.Println("Error: " + err.Error())
		return errors.New("unable to ping db...")
	}

	DB = db
	log.Println("successfuly connected with database")
	return nil
}
