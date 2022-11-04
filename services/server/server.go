package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

var (
	Server = echo.New()
)

func InitMiddleware() {
	LoggerMiddleware()
}

func RunServer(host, port string) {
	Server.Logger.Fatal(Server.Start(fmt.Sprintf("%s:%s", host, port)))
}
