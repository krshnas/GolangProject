package middlewares

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// RequestIDMiddleware ensures every request has a unique Request ID
func RequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Request().Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = uuid.New().String()
			}
			c.Set("RequestID", requestID) // Ensure it's set in context
			c.Response().Header().Set("X-Request-ID", requestID)
			return next(c)
		}
	}
}
