package todo

import entity "github.com/alifudin-a/golang-todo-app/pkg/domain/entity/todo"

type TodoRepository interface {
	Login(arg LoginParam) (*entity.User, error)
}

type LoginParam struct {
	Username string
}