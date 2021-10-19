package todo

import (
	database "github.com/alifudin-a/golang-todo-app/pkg/database/postgres"
	entity "github.com/alifudin-a/golang-todo-app/pkg/domain/entity/todo"
	query "github.com/alifudin-a/golang-todo-app/pkg/domain/query/todo"
	repository "github.com/alifudin-a/golang-todo-app/pkg/repository/todo"
)

type service struct{}

func NewTodoRepository() repository.TodoRepository {
	return &service{}
}

func (*service) Login(arg repository.LoginParam) (*entity.User, error) {
	var db = database.DB
	var user entity.User

	err := db.Get(&user, query.Login, arg.Username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}