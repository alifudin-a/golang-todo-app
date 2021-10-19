package auth

import (
	"fmt"
	entity "github.com/alifudin-a/golang-todo-app/pkg/domain/entity/todo"
	"github.com/alifudin-a/golang-todo-app/pkg/domain/helper"
	service "github.com/alifudin-a/golang-todo-app/pkg/service/todo"
	"github.com/labstack/echo/v4"
	builder "github.com/alifudin-a/golang-todo-app/pkg/domain/builder/todo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type login struct{}

// NewLoginHandler : login handler
func NewLoginHandler() *login{
	return &login{}
}

func (*login) validate (req *entity.User, c echo.Context) (err error){
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (l *login) LoginHandler(c echo.Context) (err error){
	var resp helper.Response
	var user *entity.User
	var req = new(entity.User)

	err = l.validate(req, c)
	if err != nil {
		return err
	}

	srvc := service.NewTodoRepository()

	arg := builder.LoginBuilder(req)

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