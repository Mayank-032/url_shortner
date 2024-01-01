package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"short-url/infrastructure/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMySQL() error {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.Configuration.Database.Username, config.Configuration.Database.Password, config.Configuration.Database.Host, config.Configuration.Database.Port, config.Configuration.Database.Schema)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	err = db.PingContext(ctx)
	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}

	DB = db
	log.Println("successfuly connected with database")
	return nil
}
