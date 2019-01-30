package function

import (
	"encoding/json"
	"errors"

	"github.com/dghubble/go-twitter/twitter"
	oa1 "github.com/dghubble/oauth1"
)

var (
	errValNotExists = errors.New("Consumer key/secret and Access token/secret required")
	errInvalidData  = errors.New("Invalid request data")
)

type req struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	AccessToken    string `json:"access_token"`
	AccessSecret   string `json:"access_secret"`
	Msg            string `json:"msg"`
}

func (r *req) getClient() (*twitter.Client, error) {
	if r.ConsumerKey == "" || r.ConsumerSecret == "" || r.AccessToken == "" || r.AccessSecret == "" {
		return nil, errValNotExists
	}
	config := oa1.NewConfig(r.ConsumerKey, r.ConsumerSecret)
	token := oa1.NewToken(r.AccessToken, r.AccessSecret)
	httpClient := config.Client(oa1.NoContext, token)
	return twitter.NewClient(httpClient), nil
}

// Handle a serverless request
func Handle(r []byte) string {
	var request = &req{}
	if ok := json.Unmarshal(r, request); ok != nil {
		return errInvalidData.Error()
	}
	c, err := request.getClient()
	if err != nil {
		return err.Error()
	}
	_, _, err = c.Statuses.Update(request.Msg, nil)
	if err != nil {
		return err.Error()
	}
	return "succssed"
}
