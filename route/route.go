package route

import (
	"github.com/labstack/echo"

	"go_sample/hello"
	"go_sample/serach"
    "go_sample/twitter"
    "go_sample/ob"

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
    // $ curl -X POST http://localhost:8081/api/add -H 'Content-Type: application/json' -d '{ "name": "名前" , "Description": "解説"}'
    e.POST("/api/add", ob.AddFavorite)
    // $ curl -X POST http://localhost:8081/api/find -d 'name=名前'
    // $ curl -X POST http://localhost:8081/api/find -d 'description=解説'
    // $ curl -X POST http://localhost:8081/api/find -d 'keyword=キーワード'
    e.POST("/api/find", ob.Find)
    // $ curl -X POST http://localhost:8081/api/update -d 'name=名前&description=解説　更新'
    e.POST("/api/update", ob.Update)
    e.GET("/api/get/all", ob.GetAll)
    // $ curl -X POST http://localhost:8081/api/remove -d 'name=名前'
    e.POST("/api/remove", ob.Remove)
	e.Logger.Fatal(e.Start(":8080"))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Template struct {
	templates *template.Template
}
