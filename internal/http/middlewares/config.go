package middlewares

import (
	"github.com/labstack/echo/v4/middleware"
)

type userConfig struct {
	Skipper middleware.Skipper
}

const (
	headerUserID        = "user-id"
	headerAdmin         = "admin"
	headerUserFirstName = "user-firstname"
	headerUserLastName  = "user-lastname"
	headerUserEmail     = "user-email"
	headerServiceName   = "service-name"
)
