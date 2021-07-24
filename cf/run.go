package cf

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/davidc6/cf-careers/utils"
)

const url = "https://www.cloudflare.com/en-gb/careers/jobs/?department=Engineering&location=London,%20United%20Kingdom"
const ID = "cf"
const LinksPage = "index.html"
const LinksPageFile = ID + "/" + LinksPage
const FirstPage = ID + "/page_0.html"

var links []string

func parse() {
	body := utils.StringToReader(LinksPageFile)

	loadDoc(body)
}

func parseLink(keyword string) {
	var arr []string
	
	for _, link := range links {
		end := strings.Split(link, "jobs/")
		id := strings.Split(end[1], "?")
		
		str := ID + "/" + id[0] + ".html"
		body := utils.StringToReader(str)
		
		doc, err := goquery.NewDocumentFromReader(body)
		if (err != nil) {
			log.Fatal(err)
		}
		
		s := doc.Find("#content").Text()
		
		var keywords []string
		if (keyword == "c" || keyword == "d" || keyword == "lua") {
			keywords = append(keywords, " " + keyword + ".")
			keywords = append(keywords, ", " + keyword + " ")
			keywords = append(keywords, " " + keyword + ",")
			keywords = append(keywords, "and " + keyword + ".")
			keywords = append(keywords, " " + keyword + " ")
		}

		for _, val := range keywords {
			c := strings.Contains(strings.ToLower(s), strings.ToLower(val))

			if c {
				arr = append(arr, id[0])
				break
			}
		}
	}

	if (len(arr) == 0) {
		fmt.Println("No match")
	} else {		
		for _, val := range arr {
			fmt.Println("https://boards.greenhouse.io/cloudflare/jobs/" + val + ".html")
		}
	}
}

func Run(keyword string) {
	// [1] Check if already fetched
	if (utils.DoesFileExist(LinksPageFile)) {
		fmt.Println("File exists, parsing ...")
		parse()

		// parse all links
		parseLink(strings.ToLower(keyword))

		return
	}

	// [2] File doesn't exist, we need to fetch and save
	fmt.Println("File does not exist, fetching ...")
	
	// create links page
	utils.CreateDir(ID)

	fmt.Println("Fetching jobs page ...")
	if data, err := utils.MakeRequestHeadless(url); err != nil {
		log.Fatal(err)
	} else {
		utils.SaveToDisk("files/" + LinksPageFile, strings.NewReader(data))
		fmt.Println("Saved")
		fmt.Println("Fetching jobs ...")
		parse()
		fmt.Println("Saved")
		
		// parse all links
		parseLink(keyword)
	}
}

func LinksToFetch(doc *goquery.Document) ([]string) {
	links := make([]string, 0)

	searchBy := "#jobs-list [style=\"\"]"

	doc.Find(searchBy).Each(func(i int, s *goquery.Selection) {
		link := s.Find("a").AttrOr("href", "")
		links = append(links, link)
	})

	return links
}

func loadDoc(body io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	
	if err != nil {
		log.Fatal((err))
	}

	links = LinksToFetch(doc)
	
	rand := false
	
	if (rand) {
		ch := make(chan []byte)
	
		for _, link := range links {
			go utils.MakeRequestAsync(link, ch)

			end := strings.Split(link, "jobs/")
			id := strings.Split(end[1], "?")
	
			utils.SaveToDisk("files/cf/" + id[0] + ".html", bytes.NewReader(<-ch))
		}

		return "", err
	}

	return "", nil
}
