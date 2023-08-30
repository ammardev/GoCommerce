package connections

import (
	"log"
	"time"

	"github.com/ammardev/gocommerce/internal"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func NewMySqlConnection() {
	var err error
	DB, err = sqlx.Open("mysql", buildDatasource())
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	setConnectionSettings()
}

func buildDatasource() string {
	dataSourceName := internal.GetEnv("DB_USER", "root") + ":"
	dataSourceName += internal.GetEnv("DB_PASSWORD", "") + "@tcp("
	dataSourceName += internal.GetEnv("DB_HOST", "127.0.0.1") + ":"
	dataSourceName += internal.GetEnv("DB_PORT", "3306") + ")/"
	dataSourceName += internal.GetEnv("DB_DATABASE", "")

	return dataSourceName
}

func setConnectionSettings() {
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
