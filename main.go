package main

import (
	"errors"
	"fmt"
	"os"

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

func hasBeenParsed(name string, keyword string) (bool, error) {
	switch name {
	case "cf":
		cf.Run(keyword)
		return true, nil
	default:
		return false, errors.New("parser '" + name + "' not found")
	}		
}

func main() {
	if len(os.Args) > 1 {
		name := os.Args[1] // parser
		keyword := os.Args[2] // keyword (only single keyword for now)

		if _, e := hasBeenParsed(name, keyword); e != nil {
			fmt.Println("Failed:", e)
		} else {
			fmt.Println("Succeeded!")
		}		
	} else {
		fmt.Println("Please indicate a parser that you want to use")		
	}
}
