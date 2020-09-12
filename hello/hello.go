package hello

import (
    "fmt"
	"runtime"
	"net/http"
	"github.com/carlescere/scheduler"

	"github.com/labstack/echo"
)

func SayHello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
}

func runSuccess() {
    // 5秒に1回 success!! と出力させる
    scheduler.Every(5).Seconds().Run(printSuccess)
    runtime.Goexit()
}

func printSuccess() {
    fmt.Printf("success!! \n")
}

