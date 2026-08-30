package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	/*response, err := http.Get("https://example.org")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}*/

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, _ := http.NewRequest("GET", "https://example.org", nil)
	req.Header.Set("User-Agent", "Custom user agent")

	// wykonanie żądania HTTP
	response, _ := client.Do(req)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(body))
	fmt.Println(response.Status, response.StatusCode)
	headers := response.Header
	for h, v := range headers {
		fmt.Printf("%s : %s\n", h, v)
	}

}
