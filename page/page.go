package page

import (
	"net/http"

    "github.com/labstack/echo"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func Tweet(c echo.Context) error {
    return c.Render(http.StatusOK, "tweet.html", nil)
}