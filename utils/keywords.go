package utils

import "strings"

func Keywords(keyword string) ([]string) {
	var keywords []string

	if (keyword == "c" || keyword == "d" || keyword == "lua") {
		keywords = append(keywords, " " + keyword + ".")
		keywords = append(keywords, ", " + keyword + " ")
		keywords = append(keywords, " " + keyword + ",")
		keywords = append(keywords, "and " + keyword + ".")
		keywords = append(keywords, " " + keyword + " ")
	}

	return append(keywords, keyword)
}

func RoleID(link string) (string) {
	end := strings.Split(link, "jobs/")
	id := strings.Split(end[1], "?")
	
	return id[0]
}
