package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/davidc6/cf-careers/cf"
	"github.com/davidc6/cf-careers/utils"
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

func hasBeenParsed(parserId string, keyword string, shouldRefetch bool) (bool, error) {
	utils.CreateDirIfNotExists("files")

	switch parserId {
	case "cf":
		cf.Run(keyword, shouldRefetch)
		return true, nil
	default:
		return false, errors.New("parser '" + parserId + "' not found")
	}		
}

func main() {
	if len(os.Args) > 1 {
		parserId := os.Args[1] // parser id
		keyword := os.Args[2] // keyword (only single keyword for now)
		refetch := false

		if _, e := hasBeenParsed(parserId, keyword, refetch); e != nil {
			fmt.Println("Failed:", e)
		} else {
			fmt.Println("Succeeded!")
		}
	} else {
		fmt.Println("Please supply a parser id and keyword")		
	}
}
