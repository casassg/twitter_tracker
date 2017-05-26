package main

// OAuth1
import (
	"github.com/ChimeraCoder/anaconda"
	"os"
	//"net/url"
	"net/url"
	"encoding/json"
)

func main() {
	var access_token = os.Getenv("ACCESS_TOKEN")
	var token_secret = os.Getenv("ACCESS_TOKEN_SECRET")
	var consumer_key = os.Getenv("CONSUMER_KEY")
	var consumer_secret = os.Getenv("CONSUMER_SECRET")
	//var tokens = os.Getenv("TOKENS")

	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	client := anaconda.NewTwitterApi(access_token, token_secret)

	values := url.Values{}
	values.Set("track", "Trump")
	stream := client.PublicStreamFilter(values)

	//stream := client.PublicStreamSample(nil)

	for {
		item := <-stream.C
		jsonTweet, _ := json.Marshal(item)
		println(string(jsonTweet))

	}
}
