package router

import (
	helper "github.com/alifudin-a/golang-todo-app/pkg/domain/helper"
	auth "github.com/alifudin-a/golang-todo-app/pkg/http/rest/auth"
	todo "github.com/alifudin-a/golang-todo-app/pkg/http/rest/todo"
	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// Router : app router
func Router() *echo.Echo {
	e := echo.New()

	e.Validator = &helper.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = helper.CustomReadableError

	// middlewares
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "==> METHOD=${method}, URI=${uri}, STATUS=${status}, " +
			"HOST=${host}, ERROR=${error}, LATENCY_HUMAN=${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderContentDisposition, "app-key", "user-token"},
		ExposeHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderContentDisposition, "app-key", "user-token"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	api := e.Group("/api")
	v1 := api.Group("/v1")

	// login
	v1.POST("/login", auth.NewLoginHandler().LoginHandler)

	// todo
	v1.POST("/todo", todo.NewAddTodoHandler().AddTodoHandler)
	v1.PUT("/todo", todo.NewUpdateTodoHandler().UpdateTodoHandler)
	v1.GET("/todo/:id", todo.NewGetTodoHandler().GetTodoHandler)
	v1.GET("/todo/owner/:id", todo.NewListTodoHandler().ListTodoHandler)
	v1.DELETE("/todo/:id", todo.NewDeleteTodoHandler().DeleteTodoHandler)

	e.Logger.Fatal(e.Start(":4321"))

	return e
}
