package middlewares

import (
	"github.com/go-logr/logr"
	"github.com/labstack/echo/v4"
)

// RequestLogger is a custom middleware for logging requests.
func RequestLogger(logger logr.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			log := logger.WithName("server").WithValues(
				"method", req.Method,
				"path", req.URL.Path,
			)

			log.Info("Request received")
			err := next(c)

			log = log.WithValues("status", res.Status)
			if err != nil {
				log.Error(err, "Request error")
				c.Error(err)
			}
			return err
		}
	}
}
