package todo

import (
	entity "github.com/alifudin-a/golang-todo-app/pkg/domain/entity/todo"
	"github.com/alifudin-a/golang-todo-app/pkg/domain/helper"
	repository "github.com/alifudin-a/golang-todo-app/pkg/repository/todo"
	"strings"
	"time"
)

// LoginBuilder : paramater and argument binder
func LoginBuilder(param *entity.User) repository.LoginParam {
	var arg repository.LoginParam

	username := strings.ToLower(param.Username)

	arg.Username = username

	return arg
}

func AddTodoBuilder(param *entity.Todo) repository.AddTodoParam{
	var arg repository.AddTodoParam
	t := time.Now()

	arg.Title = param.Title
	arg.Description = param.Description
	arg.OwnerID = param.OwnerID
	arg.CreatedAt = helper.NullString(t.Format(helper.LayoutTime1))

	return arg
}

func UpdateTodoBuilder(param *entity.Todo) repository.UpdateTodoParam{
	var arg repository.UpdateTodoParam
	t := time.Now()

	arg.ID = param.ID
	arg.Title = param.Title
	arg.Description = param.Description
	arg.UpdatedAt = helper.NullString(t.Format(helper.LayoutTime1))

	return arg
}