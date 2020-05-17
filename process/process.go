package process

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// URL finds all links and process them with process.URL
func URL(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

// Img finds all images and process them with process.Img
func Img(index int, element *goquery.Selection) {
	src, exists := element.Attr("src")
	if exists {
		fmt.Println(src)
	}
}

// Script finds all scripts and process them with process.Script
func Script(index int, element *goquery.Selection) {
	script, exists := element.Attr("src")
	if exists {
		fmt.Println(script)
	}
}

// Mail finds all mailto: and process them with process.Mail
func Mail(index int, element *goquery.Selection) {
	mail, exists := element.Attr("href")
	if exists && strings.HasPrefix(mail, "mailto:") {
		fmt.Println(mail)
	}
}

// Style finds all link containing rel="stylesheet" and process them with process.Style
func Style(index int, element *goquery.Selection) {
	rel, exists := element.Attr("rel")
	if exists && rel == "stylesheet" {
		style, exists := element.Attr("href")
		if exists {
			fmt.Println(style)
		}
	}
}
