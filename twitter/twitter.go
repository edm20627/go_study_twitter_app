package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
	"github.com/labstack/echo"
)

func Serach(c echo.Context) error {
	keyword := c.FormValue("keyword")
	api := ConnectTwitterApi()
	// 検索
	searchResult, _ := api.GetSearch(`"`+keyword+`"`, nil)

	tweets := make([]*Tweet, 0)

	for _, data := range searchResult.Statuses {
		tweet := new(Tweet)
		tweet.Text = data.FullText
		tweet.User = data.User.Name
		tweets = append(tweets, tweet)
	}

	return c.JSON(http.StatusOK, tweets)
}

func ConnectTwitterApi() *anaconda.TwitterApi {
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

// TwitterAccount はTwitterの認証用の情報
type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

// Tweet はツイートの情報
type Tweet struct {
	User string `json:"user"`
	Text string `json:"text"`
}

// Tweets はTweetの配列
type Tweets *[]Tweet
