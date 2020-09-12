package route

import (
    "github.com/labstack/echo"

	"go_sample/hello"
)

func Route() {
    e := echo.New()
	e.GET("/", hello.SayHello)
    e.Logger.Fatal(e.Start(":8080"))
}
