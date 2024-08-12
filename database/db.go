package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func MySQLDB(config mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal("Database open connection failed", err)
	}

	// See "Important settings" section.
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)
	// log.Println("Connected to MySQL!")
	return db, err
}
