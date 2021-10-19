package todo

import entity "github.com/alifudin-a/golang-todo-app/pkg/domain/entity/todo"

// TodoRepository : repository contract for todo service
type TodoRepository interface {
	Login(arg LoginParam) (*entity.User, error)
	Add(arg AddTodoParam) (*entity.Todo, error)
	Update(arg UpdateTodoParam) (*entity.Todo, error)
	IsExist(arg IsExistParam) (bool, error)
	IsOwnerExist(arg IsOwnerExistParam) (bool, error)
	Get(arg GetTodoParam) (*entity.Todo, error)
	List(arg  ListTodoParam) ([]entity.Todo, error)
	Delete(arg DeleteTodoParam) error
}

// LoginParam : parameter for login auth
type LoginParam struct {
	Username string
}

type AddTodoParam struct {
	entity.Todo
}

type UpdateTodoParam struct {
	entity.Todo
}

type IsExistParam struct {
	ID int
}

type GetTodoParam struct {
	ID int
}

type ListTodoParam struct {
	OwnerID int
}

type IsOwnerExistParam struct {
	OwnerID int
}

type DeleteTodoParam struct {
	ID int
}