package server

import (
	"time"

	"github.com/Mohsen-Khodabakhshi/gosta/services/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware() {
	Server.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			currentTime := time.Now()

			logs := log.LoadHttp(v.Method, v.URI, currentTime.Format("2006-01-02"), currentTime.Format("15:04:05"), v.Status)

			log.HttpOnConsole(logs)
			log.HttpToFile(logs)

			return nil
		},
	}))
}
