package main

// OAuth1
import (
	"github.com/dghubble/oauth1"
	"os"
	"bufio"
	"strings"
)

func scanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := strings.Index(string(data), "\r\n"); i >= 0 {
		// We have a full '\r\n' terminated line.
		return i + 2, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\n' {
		return data[0: len(data)-1]
	}
	return data
}

func main() {
	var access_token = os.Getenv("ACCESS_TOKEN")
	var token_secret = os.Getenv("ACCESS_TOKEN_SECRET")
	var consumer_key = os.Getenv("CONSUMER_KEY")
	var consumer_secret = os.Getenv("CONSUMER_SECRET")
	//var tokens = os.Getenv("TOKENS")

	conf := oauth1.NewConfig(consumer_key, consumer_secret)
	token := oauth1.NewToken(access_token, token_secret)
	client := conf.Client(oauth1.NoContext, token)
	stream_url := "https://stream.twitter.com/1.1/statuses/filter.json?track=Trump"

	resp, _ := client.Post(stream_url, "application/json", nil)

	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(scanLines)

	count := 0
	for scanner.Scan() {
		token := scanner.Bytes()
		if len(token) == 0 {
			// empty keep-alive
			continue
		}
		//println(string(token))
		//println("-----------------------")
		println(count)
		count += 1

	}
}
