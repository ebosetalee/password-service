package main

import (
	"log"
	"os"

	"github.com/ebosetalee/password-service.git/config"
	db "github.com/ebosetalee/password-service.git/database"
	mysqlDriver "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4"
	sqlDriver "github.com/golang-migrate/migrate/v4/database/mysql"
)

func main() {
	var env = config.Env
	db, err := db.MySQLDB(mysqlDriver.Config{
		User:                 env.DBUser,
		Passwd:               env.DBPassword,
		Addr:                 env.DBAddress,
		DBName:               env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal("Migration Database setup failed: ", err)
	}

	driver, err := sqlDriver.WithInstance(db, &sqlDriver.Config{})
	if err != nil {
		log.Fatal("Migration Driver setup failed: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://cmd/migrate/migrations", "mysql", driver)
	if err != nil {
		log.Fatal("Migration instance failed: ", err)
	}

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	// if cmd == "status" {
    //     err := yourMigrationFunctionHere() // Replace with actual function
    //     if err != nil {
    //         log.Fatalf("Error running migration: %v", err)
    //     }
    // }
}
