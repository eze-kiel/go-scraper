package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/eze-kiel/go-scraper/process"
	"github.com/fatih/color"
)

var target, contentType string

func main() {

	flag.StringVar(&target, "t", "", "target's url")
	flag.StringVar(&contentType, "type", "url", "content type that will be scraped")
	flag.Parse()

	if target == "" {
		color.Red("You must at least provide target's URL : -t <URL>")
		os.Exit(1)
	}

	// Make HTTP request
	response, err := http.Get(target)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var document *goquery.Document
	var body []byte

	if contentType == "comments" {
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal("Error reading HTTP body. ", err)
		}
	} else {
		// Create a goquery document from the HTTP response
		document, err = goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal("Error loading HTTP response body. ", err)
		}
	}

	// Process the desired content
	switch contentType {
	case "url":
		document.Find("a").Each(process.URL)
	case "img":
		document.Find("img").Each(process.Img)
	case "script":
		document.Find("script").Each(process.Script)
	case "mail":
		document.Find("a").Each(process.Mail)
	case "style":
		document.Find("link").Each(process.Style)
	case "comments":
		process.Comments(body)
	default:
		color.Red("This type is not recognised")
		os.Exit(1)
	}
}
