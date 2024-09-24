package controllers

import (
	"clean-arch/internal/domain"
	m "clean-arch/internal/http/middlewares"
	"github.com/labstack/echo/v4"
)

func parseUser(c echo.Context) *domain.User {
	if c.Get("user") == nil {
		user := m.GenerateMetadata(c, nil)
		return user
	}
	return c.Get("user").(*domain.User)
}
