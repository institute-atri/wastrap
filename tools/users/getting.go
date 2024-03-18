package users

import (
	"encoding/json"

	"github.com/institute-atri/glogger"
	"github.com/institute-atri/gnet"
)

func Getting(url string) []string {
	urlJson := url + "/wp-json/wp/v2/users"
	response := gnet.GET(urlJson)

	var users []map[string]interface{}
	err := json.Unmarshal([]byte(response.BRaw), &users)
	if err != nil {
		glogger.Fatal("Error decoding the request:", err)
	}

	var slugs []string
	for _, user := range users {
		if slug, ok := user["slug"].(string); ok {
			slugs = append(slugs, slug)
		}
	}

	return slugs
}
