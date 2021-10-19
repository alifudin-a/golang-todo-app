package todo

import (
	database "github.com/alifudin-a/golang-todo-app/pkg/database/postgres"
	entity "github.com/alifudin-a/golang-todo-app/pkg/domain/entity/todo"
	query "github.com/alifudin-a/golang-todo-app/pkg/domain/query/todo"
	repository "github.com/alifudin-a/golang-todo-app/pkg/repository/todo"
)

type service struct{}

// NewTodoRepository : repository chain for todo service
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

func (*service) Add(arg repository.AddTodoParam) (*entity.Todo, error) {
	var db = database.DB
	var todo entity.Todo

	tx := db.MustBegin()
	err := tx.QueryRowx(
		query.AddTodo,
		arg.Title,
		arg.Description,
		arg.OwnerID,
		arg.CreatedAt,
		).StructScan(&todo)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (*service) Update(arg repository.UpdateTodoParam) (*entity.Todo, error){
	var db = database.DB
	var todo entity.Todo

	tx := db.MustBegin()
	err := tx.QueryRowx(query.UpdateTodo, arg.Title, arg.Description, arg.UpdatedAt, arg.ID).StructScan(&todo)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (*service) IsExist(arg repository.IsExistParam) (bool, error){
	var db = database.DB
	var total int

	err  := db.Get(&total, query.IsExist, arg.ID)
	if err != nil {
		return false, nil
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

func (*service) Get(arg repository.GetTodoParam) (*entity.Todo, error) {
	var db = database.DB
	var todo entity.Todo

	err := db.Get(&todo, query.GetTodo, arg.ID)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (*service) List(arg  repository.ListTodoParam) ([]entity.Todo, error){
	var db = database.DB
	var todo []entity.Todo

	err := db.Select(&todo, query.ListTodo, arg.OwnerID)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (*service) IsOwnerExist(arg repository.IsOwnerExistParam) (bool, error){
	var db = database.DB
	var total int

	err := db.Get(&total, query.IsOwnerExist, arg.OwnerID)
	if err != nil {
		return false, err
	}

	if total == 0 {
		return false, err
	}

	return true, nil
}

func (*service) Delete(arg repository.DeleteTodoParam) error {
	var db = database.DB

	tx := db.MustBegin()
	_, err := tx.Exec(query.DeleteTodo, arg.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}