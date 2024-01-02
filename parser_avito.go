package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	var what_house string = "kvartira"
	var how string = "na_dlitelnyy_srok"
	var url_begin string = "https://www.avito.ru/moskva"
	var full_url string = fmt.Sprintf("%s/%s/%s-ASgBAgICAkSSA8gQ8AeQUg?cd=1", url_begin, how, what_house)
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
	c.Visit(full_url)

	fmt.Print(c)
}
