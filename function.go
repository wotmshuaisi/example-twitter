package kubeless

import (
	"errors"
	"log"
	"os"

	oa1 "github.com/dghubble/oauth1"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/kubeless/kubeless/pkg/functions"
)

var (
	errEnvNotExists = errors.New("Consumer key/secret and Access token/secret required")
)

func getClient() (*twitter.Client, error) {
	consumerKey := os.Getenv("CKEY")       // Twitter Consumer Key
	consumerSecret := os.Getenv("CSECRET") // Twitter Consumer Secret
	accessToken := os.Getenv("ATOKEN")     // Twitter Access Token
	accessSecret := os.Getenv("ASECRET")   // Twitter Access Secret
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		return nil, errEnvNotExists
	}
	config := oa1.NewConfig(consumerKey, consumerSecret)
	token := oa1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oa1.NoContext, token)
	return twitter.NewClient(httpClient), nil
}

// SendTweet serverless function use to send tweets
func SendTweet(event functions.Event, ctx functions.Context) (string, error) {
	log.Printf("%+v\n", event)
	c, err := getClient()
	if err != nil {
		return "failed", err
	}
	_, _, err = c.Statuses.Update(event.Data, nil)
	if err != nil {
		return "failed", err
	}
	return "succssed", nil
}
