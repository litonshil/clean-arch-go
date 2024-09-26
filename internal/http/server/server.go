package server

import (
	"clean-arch/client/logger"
	"clean-arch/config"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	Echo *echo.Echo
}

func (s *Server) Start(lc logger.LogClient) {
	e := s.Echo

	go func() {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.App().Port)))
	}()

	GracefulShutdown(e, lc)
}

func New() *Server {
	return &Server{Echo: echo.New()}
}

// GracefulShutdown server will gracefully shut down within 5 sec
func GracefulShutdown(e *echo.Echo, lc logger.LogClient) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	lc.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = e.Shutdown(ctx)
	lc.Info("server shutdowns gracefully")
}
