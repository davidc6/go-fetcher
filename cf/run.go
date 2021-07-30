package cf

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/davidc6/cf-careers/utils"
)

const url = "https://www.cloudflare.com/en-gb/careers/jobs/?department=Engineering&location=London,%20United%20Kingdom"
const parserID = "cf"
const mainPage = "index.html"
const mainPagePath = "/" + parserID + "/" + mainPage

func fetchAndSaveIfRequired(urls []string, shouldRefetch bool) {
	ch := make(chan []byte)
	
	for _, url := range urls {
		path := "cf/" + utils.RoleID(url) + ".html"

		if (!utils.DoesRegularFileExist(path) || shouldRefetch) {
			go utils.MakeRequestAsync(url, ch)
			utils.SaveToDisk(path, bytes.NewReader(<-ch))
		}
	}
}


func searchFilesForKeyword(urls []string, searchFor string) {
	var matchedRoles = make([]string, 0)

	for _, url := range urls {
		roleId := utils.RoleID(url)

		file := parserID + "/" + roleId + ".html"
		body := utils.StringToReader(file)

		doc, err := goquery.NewDocumentFromReader(body)

		if (err != nil) {
			log.Fatal(err)
		}

		bodyString := doc.Find("#content").Text()
		keywords := utils.Keywords(searchFor)

		for _, keyword := range keywords {
			doesContainKeyword := strings.Contains(strings.ToLower(bodyString), strings.ToLower(keyword))
			if doesContainKeyword {				
				matchedRoles = append(matchedRoles, roleId)
				break
			}
		}
	}

	if (len(urls) == 0) {
		fmt.Println("No match")
	} else {
		for _, id := range matchedRoles {
			fmt.Println("https://boards.greenhouse.io/cloudflare/jobs/" + id + ".html")
		}
	}
}

func lastSteps(keyword string, shouldRefetch bool, message string) {
	urls := extractUrls(mainPagePath)

	fetchAndSaveIfRequired(urls, shouldRefetch)

	fmt.Println(message)

	searchFilesForKeyword(urls, keyword)
}

func Run(keyword string, shouldRefetch bool) {
	// [1] Exists
	if (utils.DoesDirExist(parserID) && utils.DoesRegularFileExist(mainPagePath)) {
		fmt.Println("File exists, processing ...")
		
		lastSteps(keyword, shouldRefetch, "Main file parsed")

		return
	}

	// [2] Does not exist
	fmt.Println("File does not exist, fetching ...")
	utils.CreateDir(parserID)
	fmt.Println("Fetching jobs page ...")

	if data, err := utils.MakeRequestHeadless(url); err != nil {
		log.Fatal(err)
	} else {
		utils.SaveToDisk(mainPagePath, strings.NewReader(data))

		fmt.Println("Saved")
		fmt.Println("Fetching jobs ...")

		lastSteps(keyword, shouldRefetch, "Saved")
	}
}

func extractUrls(path string) []string {
	body := utils.StringToReader(path)
	doc, err := goquery.NewDocumentFromReader(body)
	
	if err != nil {
		log.Fatal(err)
	}

	urls := make([]string, 0)

	searchBy := "#jobs-list [style=\"\"]"

	doc.Find(searchBy).Each(func(i int, s *goquery.Selection) {
		url := s.Find("a").AttrOr("href", "")		
		urls = append(urls, url)
	})

	return urls
}
