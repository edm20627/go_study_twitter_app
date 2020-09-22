package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

const url = "http://localhost:8080/callback"

// ツイッターの認証開始
func AuthTwitter(c echo.Context) error {
	api := connectAPI()
	// 認証
	uri, _, error := api.AuthorizationURL(url)
	if error != nil {
		fmt.Println(error)
		return error
	}
	// 成功したらTwitterのログイン画面へ
	return c.Redirect(http.StatusFound, uri)
}

// ログイン後のコールバックから認証まで
func Callback(c echo.Context) error {
	token := c.QueryParam("oauth_token")
	secret := c.QueryParam("oauth_verifier")
	api := connectAPI()

	cred, _, error := api.GetCredentials(&oauth.Credentials{
		Token: token,
	}, secret)
	if error != nil {
		fmt.Println(error)
		return error
	}
	api = anaconda.NewTwitterApi(cred.Token, cred.Secret)

	sess := session.Default(c)
	sess.Set("token", cred.Token)
	sess.Set("secret", cred.Secret)
	sess.Save()

	return c.Redirect(http.StatusFound, "./tweet")
}

// HasCookie クッキーあるかどうか確認
func HasCookie(c echo.Context) error {
	cookie, _ := c.Cookie("twitter")
	if cookie == nil {
		return c.JSON(http.StatusNoContent, "no")
	}
	return c.JSON(http.StatusOK, "has data")
}

// ツイッター投稿
func PostTwitterAPI(c echo.Context) error {
	sess := session.Default(c)
	token := sess.Get("token")
	secret := sess.Get("secret")
	if token == nil || secret == nil {
		return c.JSON(http.StatusAccepted, "redirect")
	}
	api := anaconda.NewTwitterApi(token.(string), secret.(string))

	message := c.FormValue("message")
	tweet, error := api.PostTweet(message, nil)
	if error != nil {
		fmt.Println(error)
		return c.JSON(http.StatusAccepted, "redirect")
	}
	link := "https://twitter.com/" + tweet.User.IdStr + "/status/" + tweet.IdStr

	return c.JSON(http.StatusOK, link)
}

func connectAPI() *anaconda.TwitterApi {
	// Json読み込み
	raw, error := ioutil.ReadFile("./path/to/twitterAccount.json")
	if error != nil {
		fmt.Println(error.Error())
		return nil
	}

	var twitterAccount Account
	// 構造体にセット
	json.Unmarshal(raw, &twitterAccount)

	anaconda.SetConsumerKey(twitterAccount.ConsumerKey)
	anaconda.SetConsumerSecret(twitterAccount.ConsumerSecret)

	// 認証
	return anaconda.NewTwitterApi("", "")
}

type Account struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}
