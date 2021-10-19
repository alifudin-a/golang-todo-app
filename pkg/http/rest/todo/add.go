package todo

import (
	builder "github.com/alifudin-a/golang-todo-app/pkg/domain/builder/todo"
	entity "github.com/alifudin-a/golang-todo-app/pkg/domain/entity/todo"
	"github.com/alifudin-a/golang-todo-app/pkg/domain/helper"
	service "github.com/alifudin-a/golang-todo-app/pkg/service/todo"
	"github.com/labstack/echo/v4"
	"net/http"
)

type add struct{}

func NewAddTodoHandler() *add {
	return &add{}
}

func (*add) validate (req *entity.Todo, c echo.Context) (err error){
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (a *add) AddTodoHandler(c echo.Context) (err error){
	var resp helper.Response
	var todo *entity.Todo
	var req = new(entity.Todo)

	err = a.validate(req, c)
	if err != nil {
		return err
	}

	srvc := service.NewTodoRepository()

	arg := builder.AddTodoBuilder(req)

	todo, err = srvc.Add(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Successfully add todo!"
	resp.Data = map[string]interface{}{
		"todo": todo,
	}

	return c.JSON(http.StatusOK, resp)
}
