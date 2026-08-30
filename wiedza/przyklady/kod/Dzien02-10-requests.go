package main

import (
	"log"

	"github.com/asmcos/requests"
)

func main() {

	header := requests.Header{
		"User-Agent":   "super user agent",
		"moj_naglowek": "ABCDEFG",
	}
	response, err := requests.Get("http://51.91.120.89/tmp", header)
	if err != nil {
		log.Fatal(err)
	}
	//println(response.Text())

	// przekazanie danych przez POST
	data := requests.Datas{
		"key1": "value1",
		"key2": "value2",
	}
	response, _ = requests.Post("https://www.httpbin.org/post", data)
	println(response.Text())

	req := requests.Requests()
	req.SetTimeout(30)
	req.Debug = 1
	response, _ = req.Get("https://www.httpbin.org/get")
	println(response.R.Status)
	println(response.Text())

}
