package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var what_house string = "kvartiry"
	var how string = "na_dlitelnyy_srok"
	var url_begin string = "https://www.avito.ru/moskva"

	var full_url string = fmt.Sprintf("%s/%s/sdam/%s", url_begin, what_house, how)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", full_url, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Print("Failed to parse the HTML document", err)
	}

	nameHTMLElement := doc.Find("div.iva-item-title-py3i_").First()
	name := nameHTMLElement.Find("a").Text()

	priceHTMLElement := doc.Find("div.price-price-JP7qe").First()
	price := priceHTMLElement.Text()

	zalogHTMLElement := doc.Find("div.iva-item-autoParamsStep-WzfS8").First()
	zalog := zalogHTMLElement.Find("p").Text()

	addressHTMLElement := doc.Find("div.geo-root-zPwRk").First()
	address := addressHTMLElement.Find("p").First().Text()
	time_metro := addressHTMLElement.Find("p.styles-module-root_top-HYzCt").Find("span.geo-periodSection-bQIE4").Text()
	metro := addressHTMLElement.Find("p.styles-module-root_top-HYzCt").Text()
	metro = strings.ReplaceAll(metro, time_metro, "")
	descriptionHTMLElement := doc.Find("div.iva-item-descriptionStep-C0ty1").First()
	description := descriptionHTMLElement.Find("p").Text()
	fmt.Println(name)
	fmt.Println(price)
	fmt.Println(zalog)
	fmt.Println(address)
	fmt.Println(metro)
	fmt.Println(time_metro)
	fmt.Println(description)

}
