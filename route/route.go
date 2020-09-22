package route

import (
	"html/template"
    "io"

	session "github.com/ipfans/echo-session"

	"github.com/labstack/echo"

	"go_sample/app"
	"go_sample/page"
	"go_sample/replace"
	// "go_sample/hello"
	// "go_sample/serach"
	// "go_sample/twitter"
	// "go_sample/ob"
)

func Route() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e := echo.New()

	e.Static("/css", "public/views/css")
	e.Static("/js", "public/views/js")
	e.File("/header.png", "public/views/header.png")

	e.Renderer = t

	//セッションを設定
	store := session.NewCookieStore([]byte("secret-key"))
	//セッション保持時間
	store.MaxAge(86400)
	e.Use(session.Sessions("ESESSION", store))

	e.GET("/", page.Index)
	e.GET("/tweet", page.Tweet)
	e.GET("/auth", app.AuthTwitter)
	e.GET("/callback", app.Callback)
	e.POST("/check", app.HasCookie)
	e.POST("/post", app.PostTwitterAPI)
	e.POST("/replace", replace.ReplaceMessage)

	// e.GET("/", hello.SayHello)
	// e.POST("/tweet", twitter.Serach)
	// e.GET("/tweet", serach.Tweet)
	// e.GET("/tweets", serach.Tweets)
	// // $ curl -X POST http://localhost:8081/api/add -H 'Content-Type: application/json' -d '{ "name": "名前" , "Description": "解説"}'
	// e.POST("/api/add", ob.AddFavorite)
	// // $ curl -X POST http://localhost:8081/api/find -d 'name=名前'
	// // $ curl -X POST http://localhost:8081/api/find -d 'description=解説'
	// // $ curl -X POST http://localhost:8081/api/find -d 'keyword=キーワード'
	// e.POST("/api/find", ob.Find)
	// // $ curl -X POST http://localhost:8081/api/update -d 'name=名前&description=解説　更新'
	// e.POST("/api/update", ob.Update)
	// e.GET("/api/get/all", ob.GetAll)
	// // $ curl -X POST http://localhost:8081/api/remove -d 'name=名前'
	// e.POST("/api/remove", ob.Remove)

	e.Logger.Fatal(e.Start(":8080"))
	// go func() {
	// 	if err := e.Start(":8080"); err != nil {
	// 		e.Logger.Info("shutting down the server")
	// 	}
	// }()
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Template struct {
	templates *template.Template
}
