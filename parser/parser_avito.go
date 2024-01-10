package parser

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Ads struct {
	Type_ads          string
	Brief_description string
	Price             string
	Deposit           string
	Address           string
	Metro             string
	Time_to_metro     string
	Description       string
}

func parser_avito(type_ads string, full_url string) {

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

	xType := fmt.Sprintf("%T", doc)
	fmt.Println(xType)

	var adss []Ads

	doc.Find("div.iva-item-body-KLUuy").Each(func(i int, e *goquery.Selection) {
		// get info about flats
		ads := Ads{}
		ads.Type_ads = type_ads
		ads.Brief_description = e.Find("div.iva-item-title-py3i_").First().Find("a").Text()
		ads.Price = e.Find("div.price-price-JP7qe").First().Text()
		ads.Deposit = e.Find("div.iva-item-autoParamsStep-WzfS8").First().Find("p").Text()

		addressHTMLElement := e.Find("div.geo-root-zPwRk").First()
		metro_dirty := addressHTMLElement.Find("p.styles-module-root_top-HYzCt").Text()

		ads.Address = addressHTMLElement.Find("p").First().Text()
		ads.Time_to_metro = addressHTMLElement.Find("p.styles-module-root_top-HYzCt").Find("span.geo-periodSection-bQIE4").Text()
		ads.Metro = strings.ReplaceAll(metro_dirty, ads.Time_to_metro, "")

		ads.Description = e.Find("div.iva-item-descriptionStep-C0ty1").First().Find("p").Text()

		adss = append(adss, ads)
	})
	name_file := type_ads + ".json"
	file, err := os.Create(name_file)
	if err != nil {

		log.Fatalln("Failed to create the output JSON file", err)

	}
	defer file.Close()

	// convert industries to an indented JSON string

	jsonString, _ := json.MarshalIndent(adss, " ", " ")

	// write the JSON string to file

	file.Write(jsonString)

}
