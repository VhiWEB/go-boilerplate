package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"

	"github.com/joho/godotenv"
)

var Db *gorm.DB

func loadEnvVars() {
	err := godotenv.Load()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file")
	}
}

func init() {
	loadEnvVars()

	dbConnection := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	if dbConnection == "mysql" {
		// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True"
		dsn := fmt.Sprint(dbUsername, ":", dbPassword, "@tcp(", dbHost, ":", dbPort, ")/", dbName, "?charset=utf8mb4&parseTime=True")
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})

		if err != nil {
			log.Panic(err)
		}

		Db = db
	} else if dbConnection == "pgsql" {
		// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
		dsn := fmt.Sprint("host=", dbHost, " user=", dbUsername, " password=", dbPassword, " dbname=", dbName, " port=", dbPort, " sslmode=disable")
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN: dsn,
		}), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})

		if err != nil {
			log.Panic(err)
		}

		Db = db
	}
}
