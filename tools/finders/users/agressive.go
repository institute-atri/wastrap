package users

import (
	"encoding/json"
	"regexp"

	"github.com/institute-atri/glogger"
	"github.com/institute-atri/gnet"
)

// Will fetch usernames aggressively and will return a slice from each method
func Agressive(url string) [][]string  {
	glogger.Warning("Method: /author-sitemap.xml")
	authorSitemapReturn := methodAuthorSitemapXML(url)
	println("=======================================")
	glogger.Warning("Method: /wp-json/wp/v2/users")
	wpJsonReturn := methodWpJson(url)
	println("=======================================")
	glogger.Warning("Method: /?rest_route=/wp/v2/users")
	RestRouterWpReturn := methodRestRouteWp(url)
	println("=======================================")

	return [][]string{authorSitemapReturn, wpJsonReturn, RestRouterWpReturn}
}

// This method will look for users on the endpoint: /author-sitemap.xml
func methodAuthorSitemapXML(url string) []string {
	urlMethod := url + "/author-sitemap.xml"
	response := gnet.GET(urlMethod)

	search := regexp.MustCompile(`<loc>.*?/author/(.*?)</loc>`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	var authors []string
	if len(matches) == 0 {
		glogger.Danger("Nothing found")
	} else {
		seenAuthor := make(map[string]bool)
		for _, match := range matches {
			if len(match) > 1 {
				author := match[1]

				if !seenAuthor[author] {
					glogger.Done("Author:", author)
					authors = append(authors, author)
					seenAuthor[author] = true
				}
			}
		}
	}
	return authors
}

// This method will look for users on the endpoint: /wp-json/wp/v2/users
func methodWpJson(url string) []string {
	urlMethod := url + "/wp-json/wp/v2/users"
	response := gnet.GET(urlMethod)

	var users []map[string]interface{}
	err := json.Unmarshal([]byte(response.BRaw), &users)
	if err != nil {
		glogger.Danger("Nothing found")
	}

	var slugs []string
	for _, user := range users {
		if slug, ok := user["slug"].(string); ok {
			slugs = append(slugs, slug)
		}
	}
	
	for _, slug := range slugs{
		glogger.Done("User:", slug)
	}

	return slugs
}

// This method will look for users on the endpoint: /?rest_route=/wp/v2/users
func methodRestRouteWp(url string) []string {
	urlMethod := url + "/?rest_route=/wp/v2/users"
	response := gnet.GET(urlMethod)

	var users []map[string]interface{}
	err := json.Unmarshal([]byte(response.BRaw), &users)
	if err != nil {
		glogger.Danger("Nothing found")
	}

	var slugs []string
	for _, user := range users {
		if slug, ok := user["slug"].(string); ok {
			slugs = append(slugs, slug)
		}
	}
	
	for _, slug := range slugs{
		glogger.Done("User:", slug)
	}

	return slugs
}