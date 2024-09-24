package middlewares

import (
	"bytes"
	"clean-arch/config"
	"clean-arch/internal/domain"
	"clean-arch/utils"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"strconv"
)

func authorizeUser(config userConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultSkipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			headers := c.Request().Header

			id, _ := strconv.Atoi(headers.Get(headerUserID))
			isAdmin, _ := strconv.ParseBool(headers.Get(headerAdmin))

			user := &domain.User{
				ID:      id,
				IsAdmin: isAdmin,
				Profile: domain.Profile{
					FirstName: headers.Get(headerUserFirstName),
					LastName:  headers.Get(headerUserLastName),
					Email:     headers.Get(headerUserEmail),
				},
			}

			c.Set("user", user)

			return next(c)
		}
	}
}

// BindBody binds request body contents to bindable object
func BindBody(c echo.Context, i interface{}) error {
	// read origin body bytes
	var bodyBytes []byte
	if !utils.IsEmpty(c.Request().Body) {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
		// write back to request body
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// parse json data
		err := json.Unmarshal(bodyBytes, &i)
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateMetadata(c echo.Context, user *domain.User) *domain.User {
	if user == nil {
		user = &domain.User{}
	}

	var body interface{}
	_ = BindBody(c, &body)
	appKey := c.Request().Header.Get(config.App().AppKeyHeader)
	if appKey != "" {
		appKey = "internal call (app key provided)"
	}
	serviceName := c.Request().Header.Get(headerServiceName)
	// metadata will be passed as slack logger metadata
	metadata := domain.Meta{
		Method:      c.Request().Method,
		URI:         c.Request().RequestURI,
		ServiceName: &serviceName,
		AppKey:      &appKey,
		Profile: domain.Profile{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
		Payload: body,
	}
	user.Metadata = metadata
	return user
}
