package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func fetchUrl(url string, wg *sync.WaitGroup, ch chan string) {

	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("download error %s : %v", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Read error %s : %v", url, err)
		return
	}

	ch <- fmt.Sprintf("Content from %s : %s", url, string(body))

}

func main() {

	urls := []string{
		"https://example.com",
		"https://example.org",
		"https://example.net",
		"https://example.io",
	}

	var wg sync.WaitGroup
	ch := make(chan string)
	for _, url := range urls {
		wg.Add(1)
		go fetchUrl(url, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

}
