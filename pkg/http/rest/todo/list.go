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

type list struct{}

func NewListTodoHandler() *list {
	return &list{}
}

func (*list) ListTodoHandler(c echo.Context) (err error){
	var resp helper.Response
	var todo []entity.Todo

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	srvc := service.NewTodoRepository()

	arg1 := repository.IsOwnerExistParam{
		OwnerID: id,
	}

	exist, err := srvc.IsOwnerExist(arg1)
	if !exist {
		resp.Code = http.StatusNotFound
		resp.Message = "Owner is not exist!"
		return c.JSON(http.StatusNotFound, resp)
	}

	arg2 := repository.ListTodoParam{
		OwnerID: id,
	}

	todo, err = srvc.List(arg2)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Successfully showing all todo!"
	resp.Data = map[string]interface{}{
		"todo": todo,
	}

	return c.JSON(http.StatusOK, resp)
}
