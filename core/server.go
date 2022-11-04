package core

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

var (
	Server = echo.New()
)

func RunServer(host, port string) {
	Server.Logger.Fatal(Server.Start(fmt.Sprintf("%s:%s", host, port)))
}
