package main

import (
    "html/template"
    "net/http"
    "github.com/labstack/echo"
    "io"

    "encoding/json"
    "fmt"
    "io/ioutil"
    "github.com/ChimeraCoder/anaconda"
)

func main() {
    e := echo.New()
    t := &Template{
        templates: template.Must(template.ParseGlob("public/views/*.html")),
    }
    e.Renderer = t
    e.GET("/tweet", tweet)
    e.GET("/tweets", tweets)


    e.Logger.Fatal(e.Start(":8080"))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func tweet(c echo.Context) error {
    value := c.QueryParam("value")
    api := connectTwitterApi()
    //検索
    searchResult, _ := api.GetSearch(`"` + value + `"`, nil)
    tweet := new(TweetTempete)
    for _, data := range searchResult.Statuses {
        tweet.Text = data.FullText
        tweet.User = data.User.Name
        tweet.Id = data.User.IdStr
        tweet.ScreenName = data.User.ScreenName
        tweet.Date = data.CreatedAt
        tweet.TweetId = data.IdStr
        break
    }
    return c.Render(http.StatusOK, "tweet.html", tweet)
}

func tweets(c echo.Context) error {
    value := c.QueryParam("value")
    api := connectTwitterApi()
    searchResult, _ := api.GetSearch(`"` + value + `"`, nil)
    tweets := make([]*TweetTempete, 0)
    for _, data := range searchResult.Statuses {
        tweet := new(TweetTempete)
        tweet.Text = data.FullText
        tweet.User = data.User.Name
        tweet.Id = data.User.IdStr
        tweet.ScreenName = data.User.ScreenName
        tweet.Date = data.CreatedAt
        tweet.TweetId = data.IdStr
        tweets = append(tweets, tweet)
    }
    return c.Render(http.StatusOK, "tweets.html", tweets)
}

type Template struct {
    templates *template.Template
}

// TweetTempete はツイートの情報
type TweetTempete struct {
    User string `json:"user"`
    Text string `json:"text"`
    ScreenName string `json:"screenName"`
    Id string `json:"id"`
    Date string `json:"date"`
    TweetId string `json:"tweetId"`
}

// あとでimportする
func connectTwitterApi() *anaconda.TwitterApi {
    // Json読み込み
    raw, error := ioutil.ReadFile("path/to/twitterAccount.json")
    if error != nil {
        fmt.Println(error.Error())
        return nil
    }

    var twitterAccount TwitterAccount
    // 構造体にセット
    json.Unmarshal(raw, &twitterAccount)

    // 認証
    return anaconda.NewTwitterApiWithCredentials(twitterAccount.AccessToken, twitterAccount.AccessTokenSecret, twitterAccount.ConsumerKey, twitterAccount.ConsumerSecret)
}

type TwitterAccount struct {
    AccessToken       string `json:"accessToken"`
    AccessTokenSecret string `json:"accessTokenSecret"`
    ConsumerKey       string `json:"consumerKey"`
    ConsumerSecret    string `json:"consumerSecret"`
}