package api

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/krishna/task-management/internal/api/middlewares"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type HTTPServer struct {
	server *http.Server
	echo   *echo.Echo
	logger logr.Logger
}

func NewHTTPServer(logger logr.Logger) *HTTPServer {
	e := echo.New()

	// Middleware
	e.Use(middlewares.RequestLogger(logger))
	e.Use(middlewares.RequestIDMiddleware())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo with Fx and Zap Logger!")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: e,
	}

	return &HTTPServer{
		server: srv,
		echo:   e,
		logger: logger,
	}
}

func RegisterServerLifecycle(lc fx.Lifecycle, srv *HTTPServer) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.server.Addr)
			if err != nil {
				return err
			}
			srv.logger.Info(fmt.Sprintf("Starting HTTP server at %s", srv.server.Addr))
			go func() {
				if err := srv.server.Serve(ln); err != nil && err != http.ErrServerClosed {
					srv.logger.Error(err, "Server failed")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.logger.Info("Shutting down server")
			return srv.server.Shutdown(ctx)
		},
	})
}
