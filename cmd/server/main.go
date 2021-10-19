package main

import (
	database "github.com/alifudin-a/golang-todo-app/pkg/database/postgres"
	router "github.com/alifudin-a/golang-todo-app/pkg/router"
	"github.com/joho/godotenv"
	"log"
)

func init(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln("An error occured while loading .env file => [", err, "]")
	}
}

func main() {
	database.OpenDB()
	router.Router()
}