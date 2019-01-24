package kubeless

import (
	"log"
	"os"
	"testing"

	"github.com/kubeless/kubeless/pkg/functions"
)

func TestSendTweet(t *testing.T) {
	os.Setenv("http_proxy", "")
	os.Setenv("https_proxy", "")
	os.Setenv("CKEY", "")
	os.Setenv("CSECRET", "")
	os.Setenv("ATOKEN", "")
	os.Setenv("ASECRET", "")

	defer func() {
		os.Unsetenv("http_proxy")
		os.Unsetenv("https_proxy")
		os.Unsetenv("CKEY")
		os.Unsetenv("CSECRET")
		os.Unsetenv("ATOKEN")
		os.Unsetenv("ASECRET")
	}()

	var e = functions.Event{
		Data: "Test Twitter",
	}

	_, err := SentTweet(e, functions.Context{})
	if err != nil {
		log.Println(err)
		t.Fatal()
	}

}
