package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupLogger configures the logging middleware with a custom format
func SetupLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"status":${status}},"time":"${time_rfc3339}", "method":"${method}", "ip":"${remote_ip}", "uri":"${uri}", "error":"${error}", "latency":"${latency}", "latency_human":"${latency_human}", "bytes_in":${bytes_in}, "bytes_out":${bytes_out}` + "\n",
	})
}
