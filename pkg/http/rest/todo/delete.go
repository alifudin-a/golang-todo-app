package todo

import (
	"github.com/alifudin-a/golang-todo-app/pkg/domain/helper"
	repository "github.com/alifudin-a/golang-todo-app/pkg/repository/todo"
	service "github.com/alifudin-a/golang-todo-app/pkg/service/todo"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type delete struct{}

func NewDeleteTodoHandler() *delete {
	return &delete{}
}

func (*delete) DeleteTodoHandler(c echo.Context) (err error){
	var resp helper.Response

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	srvc := service.NewTodoRepository()

	arg1 := repository.IsExistParam{
		ID: id,
	}

	exist, err := srvc.IsExist(arg1)
	if !exist {
		resp.Code = http.StatusNotFound
		resp.Message = "Todo is not exist!"
		return c.JSON(http.StatusNotFound, resp)
	}

	arg2 := repository.DeleteTodoParam{
		ID: id,
	}

	err = srvc.Delete(arg2)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Successfully delete todo!"

	return c.JSON(http.StatusOK, resp)
}
