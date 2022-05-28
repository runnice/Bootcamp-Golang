package main

import (
	"fmt"
	"net/http"
)

func checkStatus(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Ocorreu um erro enquanto acessava a url: %s", url)
		ch <- fmt.Sprintf("error %s", err)
		return
	}
	defer resp.Body.Close()
	u := fmt.Sprintf("url: %s - status: %s", url, resp.Status)
	ch <- u
}

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.bing.com",
		"http://www.yahoo.com",
		"http://www.facebook.com",
		"http://www.youtube.com",
		"http://www.amazon.com",
		"http://www.wikipedia.com",
		"http://www.twitter.com",
	}

	ch := make(chan string)

	for _, url := range urls {
		go checkStatus(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

}
