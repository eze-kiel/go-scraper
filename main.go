package main

import (
	"flag"
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

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Process the desired content
	switch contentType {
	case "url":
		// Find all links and process them with process.URL
		document.Find("a").Each(process.URL)
	case "img":
		// Find all images and process them with process.Img
		document.Find("img").Each(process.Img)
	case "script":
		// Find all scripts and process them with process.Script
		document.Find("script").Each(process.Script)
	case "mail":
		// Find all mailto: and process them with process.Mail
		document.Find("a").Each(process.Mail)
	case "style":
		// Find all link containing rel="stylesheet" and process them with process.Style
		document.Find("link").Each(process.Style)
	default:
		color.Red("This type is not recognised")
		os.Exit(1)
	}
}
