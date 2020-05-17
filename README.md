# go-scraper

A web scraper bot in Go.

## Usage
```
simplest usage :
go-scraper -t <URL>
```

## Options
```
-t <URL>    : URL to scrap
-type <TYPE>: content type that will be scraped. Types are detailed in the Types section
```

## Types
`url` : default value. Scrap all the links on the page

`img` : scrap all the links to images

`script` : scrap all the links to scripts

`mail` : scrap all the mailto: links

`style` : scrap all stylesheets imported