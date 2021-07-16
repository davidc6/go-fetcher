package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/davidc6/cf-careers/cf"
)

func getLinks(doc *goquery.Document) ([]string) {
	links := make([]string, 0)

	// searchBy := "#jobs-list [style=""]"
	searchBy := ".row .title"
	doc.Find(searchBy).Each(func (i int, s *goquery.Selection) {
		link := s.AttrOr("href", "")
		links = append(links, "https://webscraper.io" + link)
	})

	return links
}

func main() {
	// Todo: take command line argument
	name := "cf"

	switch name {
	case "cf":
		cf.Run()
	default:
	}
}
