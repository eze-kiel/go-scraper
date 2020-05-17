package process

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func URL(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

func Img(index int, element *goquery.Selection) {
	src, exists := element.Attr("src")
	if exists {
		fmt.Println(src)
	}
}

func Script(index int, element *goquery.Selection) {
	script, exists := element.Attr("src")
	if exists {
		fmt.Println(script)
	}
}

func Mail(index int, element *goquery.Selection) {
	mail, exists := element.Attr("href")
	if exists && strings.HasPrefix(mail, "mailto:") {
		fmt.Println(mail)
	}
}

func Style(index int, element *goquery.Selection) {
	rel, exists := element.Attr("rel")
	if exists && rel == "stylesheet" {
		style, exists := element.Attr("href")
		if exists {
			fmt.Println(style)
		}
	}
}
