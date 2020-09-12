package serach

import (
	"github.com/labstack/echo"
	"go_sample/twitter"
	"net/http"
)

func Tweet(c echo.Context) error {
	value := c.QueryParam("value")
	api := twitter.ConnectTwitterApi()
	//検索
	searchResult, _ := api.GetSearch(`"`+value+`"`, nil)
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

func Tweets(c echo.Context) error {
	value := c.QueryParam("value")
	api := twitter.ConnectTwitterApi()
	searchResult, _ := api.GetSearch(`"`+value+`"`, nil)
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

// TweetTempete はツイートの情報
type TweetTempete struct {
	User       string `json:"user"`
	Text       string `json:"text"`
	ScreenName string `json:"screenName"`
	Id         string `json:"id"`
	Date       string `json:"date"`
	TweetId    string `json:"tweetId"`
}
