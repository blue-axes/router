package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

type (
	Config struct {
	}
	Server struct {
		cfg     Config
		httpSvr *echo.Echo
	}
)

func NewServer(cfg Config) *Server {
	e := echo.New()
	s := &Server{
		cfg:     cfg,
		httpSvr: e,
	}

	e.Use(middleware.Logger())

	InitRoute(e)

	return s
}

func (s Server) Serve() error {
	var addr = fmt.Sprintf("%s:%d")
	return s.httpSvr.Start(addr)
}

func (s Server) Shutdown() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	return s.httpSvr.Shutdown(ctx)
}
