package middlewares

import (
	"clean-arch/config"
	"clean-arch/utils/monitor"
	"clean-arch/utils/msgutil"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

const (
	EchoLogFormat     = `time: ${time_custom} || Remote IP: ${remote_ip} || ${method} ${host}${uri} || status: ${status} || latency: ${latency_human} || bytes: in ${bytes_in}B out ${bytes_out}B || "${user_agent}"` + "\n"
	EchoLogTimeFormat = "2006-01-02T15:04:05.00"
)

func Init(e *echo.Echo) {
	var (
		metricsPath    = "/metrics"
		apiDocPath     = "/docs"
		swaggerDocPath = "/swagger.yaml"
	)

	e.Pre(m.RemoveTrailingSlash())
	e.Use(m.LoggerWithConfig(m.LoggerConfig{
		Format:           EchoLogFormat,
		CustomTimeFormat: EchoLogTimeFormat,
	}))
	e.Use(m.CORS())
	e.Use(m.Secure())
	e.Use(m.Recover())

	e.Use(m.GzipWithConfig(m.GzipConfig{
		Skipper: func(context echo.Context) bool {
			return context.Request().RequestURI == metricsPath
		},
		Level: 5,
	}))

	e.Use(authorizeUser(userConfig{
		Skipper: func(context echo.Context) bool {
			skipList := []string{
				apiDocPath,
				metricsPath,
				swaggerDocPath,
			}
			for _, skip := range skipList {
				if strings.HasPrefix(context.Request().URL.Path, skip) {
					return true
				}
			}
			return false
		},
	}))

	monitor.NewEchoPrometheusClient(e, &metricsPath)
}

func CheckAppKey() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			appKey := c.Request().Header.Get(config.App().AppKeyHeader)
			if appKey != config.Get().App.AppKey {
				return c.JSON(http.StatusForbidden, msgutil.ForbiddenResponseMsg())
			}
			return next(c)
		}
	}
}
