package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	url := "https://www.domiporta.pl/nieruchomosci/sprzedam-mieszkanie-trzypokojowe-warszawa-cybernetyki-62m2/154180190"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	contact := doc.Find("Div[class='agent__phone agent__phone_normal details-databox__phone-number js-details-phone-number']")
	contact.Each(func(index int, item *goquery.Selection) {
		phone, _ := item.Attr("data-tel")
		phone = strings.ReplaceAll(phone, " ", "")
		fmt.Println(phone)
	})

}
