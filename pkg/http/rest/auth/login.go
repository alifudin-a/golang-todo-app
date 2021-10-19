package auth

import (
	"fmt"
	entity "github.com/alifudin-a/golang-todo-app/pkg/domain/entity/todo"
	"github.com/alifudin-a/golang-todo-app/pkg/domain/helper"
	repository "github.com/alifudin-a/golang-todo-app/pkg/repository/todo"
	service "github.com/alifudin-a/golang-todo-app/pkg/service/todo"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type login struct{}

func NewLoginHandler() *login{
	return &login{}
}

func (*login) LoginHandler(c echo.Context) (err error){
	var resp helper.Response
	var user *entity.User
	var req = new(entity.User)

	if err = c.Bind(req); err != nil {
		return err
	}

	srvc := service.NewTodoRepository()

	arg := repository.LoginParam{
		Username: req.Username,
	}

	user, err = srvc.Login(arg)
	if err != nil {
		resp.Code = http.StatusUnauthorized
		resp.Message = "Login Failed! Please check your username and password!"
		return c.JSON(http.StatusUnauthorized, resp)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		resp.Code = http.StatusUnauthorized
		resp.Message = "Login Failed! Please check your username and password!"
		return c.JSON(http.StatusUnauthorized, resp)
	}

	msg := fmt.Sprintf("Login Success! Welcome %s!", user.FullName)

	resp.Code = http.StatusOK
	resp.Message = msg

	return c.JSON(http.StatusOK, resp)
}