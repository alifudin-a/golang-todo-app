package todo

import (
	database "github.com/alifudin-a/golang-todo-app/pkg/database/postgres"
	entity "github.com/alifudin-a/golang-todo-app/pkg/domain/entity/todo"
	repository "github.com/alifudin-a/golang-todo-app/pkg/repository/todo"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestService_Login(t *testing.T) {

	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Println("An error occured while loading .env file => [", err, "]")
	}
	database.OpenDB()

	var user *entity.User

	srvc := NewTodoRepository()

	arg := repository.LoginParam{
		Username: "valstrax",
	}

	user, err = srvc.Login(arg)
	if err != nil {
		log.Println("error: ", err)
	}

	log.Printf("Halo %s", user.Username)
}
