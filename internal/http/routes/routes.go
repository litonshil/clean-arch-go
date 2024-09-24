package routes

import (
	"clean-arch/internal/http/controllers"
	"clean-arch/internal/http/middlewares"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	echo           *echo.Echo
	authController *controllers.AuthController
	userController *controllers.UserController
}

func New(
	e *echo.Echo,
	authController *controllers.AuthController,
	userController *controllers.UserController,
) *Routes {
	return &Routes{
		echo:           e,
		authController: authController,
		userController: userController,
	}
}

func (r *Routes) Init() {
	e := r.echo
	middlewares.Init(e)

	g := e.Group("/v1")
	g.POST("/user", r.userController.CreateUser)
	g.POST("/login", r.authController.Login)

}
