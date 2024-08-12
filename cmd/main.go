package main

import (
	"database/sql"
	"log"

	"github.com/ebosetalee/password-service.git/cmd/api"
	"github.com/ebosetalee/password-service.git/config"
	db "github.com/ebosetalee/password-service.git/database"
	"github.com/go-sql-driver/mysql"
)

func main() {
	var port = ":4444"

	var env = config.Env

	db, err := db.MySQLDB(mysql.Config{
		User:                 env.DBUser,
		Passwd:               env.DBPassword,
		Addr:                 env.DBAddress,
		DBName:               env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal("Database setup failed:", err)
	}

	initDatabase(db)

	srv := api.NewAPIServer(port, db)
	if err := srv.Run(); err != nil {
		log.Fatal("server connection failed:", err)
	}
}
func initDatabase(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Connected to MySQL!")
}
