package main

import (
	"github.com/go-logr/logr"
	"github.com/krishna/task-management/internal/api"
	"github.com/krishna/task-management/pkg/utils/logger"
	"go.uber.org/fx"
)

func main() {
	logger.InitializeLogger()
	log := logger.SetLoggerName("app-start")

	app := fx.New(
		fx.Provide(
			func() logr.Logger {
				return logger.Logger
			},
			api.NewHTTPServer,
		),
		fx.Invoke(api.RegisterServerLifecycle),
	)

	app.Run()
	log.Info("Application stopped")
}
