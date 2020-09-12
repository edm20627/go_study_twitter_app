package route

import (
	"github.com/labstack/echo"

	"go_sample/hello"
	"go_sample/serach"
	"go_sample/twitter"

	"html/template"
	"io"
)

func Route() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	e.GET("/", hello.SayHello)
	e.POST("/tweet", twitter.Serach)
	e.GET("/tweet", serach.Tweet)
	e.GET("/tweets", serach.Tweets)
	e.Logger.Fatal(e.Start(":8080"))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Template struct {
	templates *template.Template
}
