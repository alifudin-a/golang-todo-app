package todo

import (
	builder "github.com/alifudin-a/golang-todo-app/pkg/domain/builder/todo"
	entity "github.com/alifudin-a/golang-todo-app/pkg/domain/entity/todo"
	"github.com/alifudin-a/golang-todo-app/pkg/domain/helper"
	repository "github.com/alifudin-a/golang-todo-app/pkg/repository/todo"
	service "github.com/alifudin-a/golang-todo-app/pkg/service/todo"
	"github.com/labstack/echo/v4"
	"net/http"
)

type update struct{}

func NewUpdateTodoHandler() *update{
	return &update{}
}

func (*update) validate(req *entity.Todo, c echo.Context) (err error) {
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (u *update) UpdateTodoHandler(c echo.Context) (err error){
	var resp helper.Response
	var todo *entity.Todo
	var req = new(entity.Todo)

	err = u.validate(req, c)
	if err != nil {
		return err
	}

	srvc := service.NewTodoRepository()

	arg := builder.UpdateTodoBuilder(req)

	arg2 := repository.IsExistParam{
		ID: req.ID,
	}

	exist, err := srvc.IsExist(arg2)
	if !exist {
		resp.Code = http.StatusNotFound
		resp.Message = "The selected todo is not exist!"
		return c.JSON(http.StatusNotFound, resp)
	}

	todo, err = srvc.Update(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Successfully update todo!"
	resp.Data = map[string]interface{}{
		"todo": todo,
	}

	return c.JSON(http.StatusOK, resp)
}
