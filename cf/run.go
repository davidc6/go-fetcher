package cf

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/davidc6/cf-careers/utils"
)

const url = "https://www.cloudflare.com/en-gb/careers/jobs/?department=Engineering&location=London,%20United%20Kingdom"
const id = "cf"

func Run() {
	// check if we need to fetch again
	if (utils.DoesFileExist(id + "/index.html")) {
		fmt.Println("file exists, parsing ...")

		data := utils.File("cf/index.html")

		loadDoc(data)

		return
	}

	// we need to fetch!
	fmt.Println("file does not exist, fetching ...")
	
	utils.CreateDir("cf")
	data, err := utils.MakeRequest(url)
	
	if (err != nil) {
		log.Fatal(err)
	}
	
	utils.SaveToDisk("cf/index.html", data)
}

func CloudflareLinks(doc *goquery.Document) ([]string) {
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
	
	links := CloudflareLinks(doc)
	
	ch := make(chan []byte)
	for _, v := range links {
		go utils.MakeRequestCon(v, ch)
	}
	
	for i, _ := range links {
		utils.SaveToDisk("files/cf/page_" + strconv.Itoa(i) + ".html", bytes.NewReader(<-ch))
	}
	
	return "", err
}
