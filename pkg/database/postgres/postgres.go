package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// DB : Global variable for sqlx.DB wrapper
var DB *sqlx.DB

// OpenDB : Open database connection
func OpenDB(){
	var (
		dbhost = os.Getenv("DB_HOST")
		dbport = os.Getenv("DB_PORT")
		dbuser = os.Getenv("DB_USER")
		dbpass = os.Getenv("DB_PASS")
		dbname = os.Getenv("DB_NAME")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbhost, dbport, dbuser, dbpass, dbname)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err = db.Ping(); err != nil {
		log.Panicln("An error occured while connectiong to database => [", err,"]")
	} else {
		log.Println("Successfully connected to database todo-app!")
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Minute * 10)

	DB = db
}
