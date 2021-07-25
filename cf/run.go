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
const linksPage = "index.html"
const linksPageFile = parserID + "/" + linksPage

func parseCareers() []string {
	body := utils.StringToReader(linksPageFile)
	doc, err := goquery.NewDocumentFromReader(body)
	
	if err != nil {
		log.Fatal(err)
	}
	
	rand := false
	
	links := LinksToFetch(doc)

	if (rand) {
		ch := make(chan []byte)
	
		for _, link := range links {
			go utils.MakeRequestAsync(link, ch)

			roleId := utils.RoleID(link)
	
			utils.SaveToDisk("files/cf/" + roleId + ".html", bytes.NewReader(<-ch))
		}

		return links
	}

	return links
}

func parseRole(roles []string, searchFor string) {
	var matchedRoles = make([]string, 0)

	for _, link := range roles {
		roleId := utils.RoleID(link)

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

	if (len(roles) == 0) {
		fmt.Println("No match")
	} else {
		for _, id := range matchedRoles {
			fmt.Println("https://boards.greenhouse.io/cloudflare/jobs/" + id + ".html")
		}
	}
}

func Run(keyword string) {
	if (utils.DoesFileExist(linksPageFile)) {
		fmt.Println("File exists, parsing ...")
		roles := parseCareers()
		fmt.Println("Main file parsed.")
		parseRole(roles, strings.ToLower(keyword))
		return
	}

	fmt.Println("File does not exist, fetching ...")
	
	// create links page
	utils.CreateDir(parserID)

	fmt.Println("Fetching jobs page ...")
	if data, err := utils.MakeRequestHeadless(url); err != nil {
		log.Fatal(err)
	} else {
		utils.SaveToDisk("files/" + linksPageFile, strings.NewReader(data))
		fmt.Println("Saved")
		fmt.Println("Fetching jobs ...")
		roles := parseCareers()
		fmt.Println("Saved")		
		parseRole(roles, keyword)
	}
}

func LinksToFetch(doc *goquery.Document) []string {
	links := make([]string, 0)

	searchBy := "#jobs-list [style=\"\"]"

	doc.Find(searchBy).Each(func(i int, s *goquery.Selection) {
		link := s.Find("a").AttrOr("href", "")
		links = append(links, link)
	})

	return links
}
