package controllers

import (
	domain2 "clean-arch/internal/domain"
	"clean-arch/utils/consts"
	"context"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	baseCtx context.Context
	useruc  domain2.UserUseCase
}

func NewUserController(
	baseCtx context.Context,
	useruc domain2.UserUseCase,
) *UserController {
	return &UserController{
		baseCtx: baseCtx,
		useruc:  useruc,
	}
}
func (cntlr *UserController) CreateUser(c echo.Context) error {
	ctx := domain2.ContextWithValue(cntlr.baseCtx, consts.ContextKeyUser, parseUser(c))
	_ = ctx
	//var err error

	return nil
	//return c.JSON(http.StatusOK, res)
}
