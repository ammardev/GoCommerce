package connections

import (
	"database/sql"
	"log"
	"time"

	"github.com/ammardev/ecommerce-playground/app"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func NewMySqlConnection() {
	DB, err := sql.Open("mysql", buildDatasource())
	if err != nil {
		log.Fatal(err)
	}

	defer DB.Close()

	setConnectionSettings()
}

func buildDatasource() string {
	dataSourceName := app.GetEnv("DB_USER", "root") + ":"
	dataSourceName += app.GetEnv("DB_PASSWORD", "") + "@tcp("
	dataSourceName += app.GetEnv("DB_HOST", "127.0.0.1") + ":"
	dataSourceName += app.GetEnv("DB_PORT", "3306") + ")/"
	dataSourceName += app.GetEnv("DB_DATABASE", "")

	return dataSourceName
}

func setConnectionSettings() {
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
