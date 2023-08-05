package connections

import (
	"database/sql"
	"log"
	"time"

	"github.com/ammardev/ecommerce-playground/app"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySqlConnection() {
	dataSourceName := app.GetEnv("DB_USER", "root") + ":"
	dataSourceName += app.GetEnv("DB_PASSWORD", "") + "@tcp("
	dataSourceName += app.GetEnv("DB_HOST", "127.0.0.1") + ":"
	dataSourceName += app.GetEnv("DB_PORT", "3306") + ")/"
	dataSourceName += app.GetEnv("DB_DATABASE", "")

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
