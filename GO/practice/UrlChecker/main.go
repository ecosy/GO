package main

import (
	"errors"
	"fmt"
	"net/http"
)

type resultRequest struct {
	url    string
	status string
}

var errResquestFailed = errors.New("Requested Failed")

func main() {
	results := make(map[string]string)
	ch := make(chan resultRequest)

	urls := []string{
		"https://www.airbnb.co.kr/",
		"https://www.google.com/",
		"https://www.reddit.com/",
		"https://www.amazon.com/",
		"https://www.naver.com/",
		"https://www.facebook.com/",
	}

	for _, url := range urls {
		go hitURL(url, ch)
	}

	// for _c := range ch {
	// fmt.Println(_c)
	// }

	for i := 0; i < len(urls); i++ {
		result := <-ch
		results[result.url] = result.status
	}

	for k, v := range results {
		fmt.Println(k, v)
	}
}

func hitURL(url string, ch chan<- resultRequest) {
	fmt.Println("checking url : ", url)

	resp, err := http.Get(url)
	status := "OK"

	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	ch <- resultRequest{url: url, status: status}
}
