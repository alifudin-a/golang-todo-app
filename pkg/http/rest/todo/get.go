package todo

import (
	entity "github.com/alifudin-a/golang-todo-app/pkg/domain/entity/todo"
	"github.com/alifudin-a/golang-todo-app/pkg/domain/helper"
	repository "github.com/alifudin-a/golang-todo-app/pkg/repository/todo"
	service "github.com/alifudin-a/golang-todo-app/pkg/service/todo"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type get struct{}

func NewGetTodoHandler() *get {
	return &get{}
}

func (*get) GetTodoHandler(c echo.Context) (err error){
	var resp helper.Response
	var todo *entity.Todo

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	srvc := service.NewTodoRepository()

	arg := repository.GetTodoParam{
		ID: id,
	}

	todo, err = srvc.Get(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Successfully get todo!"
	resp.Data = map[string]interface{}{
		"todo": todo,
	}

	return c.JSON(http.StatusOK, resp)
}
