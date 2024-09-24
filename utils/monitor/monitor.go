package monitor

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

const (
	promSubsystemName = "cleanarch"
)

// NewEchoPrometheusClient adds metricsPath to echo server as middleware
func NewEchoPrometheusClient(e *echo.Echo, metricsPath *string) {
	prom := echoprometheus.NewMiddleware(promSubsystemName)

	// Scrape metrics from Main Server
	e.Use(prom)

	if metricsPath != nil && *metricsPath != "" {
		e.GET(*metricsPath, echoprometheus.NewHandler())
	} else {
		// Default path if metricsPath is not provided
		e.GET("/metrics", echoprometheus.NewHandler())
	}
}
