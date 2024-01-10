package parser

import (
	"fmt"
)

func Room() {
	var what_house string = "komnaty"
	var how string = "na_dlitelnyy_srok"
	var url_begin string = "https://www.avito.ru/moskva"

	var full_url string = fmt.Sprintf("%s/%s/sdam/%s", url_begin, what_house, how)

	parser_avito("Room", full_url)
}
