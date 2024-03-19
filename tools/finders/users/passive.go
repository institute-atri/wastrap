package users

import (
	"regexp"
	"strings"

	"github.com/institute-atri/glogger"
	"github.com/institute-atri/gnet"
)

// It will search for usernames passively and return a slice containing the users
func Passive(url string) []string {
	urlFeed := url + "/feed"

	response := gnet.GET(urlFeed)

	search := regexp.MustCompile(`<dc:creator><!\[CDATA\[(.+?)\]\]></dc:creator>`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	var creators []string
	if len(matches) == 0 {
		glogger.Danger("Nothing found")
	} else {
		seenCreators := make(map[string]bool)
		for _, match := range matches {
			creator := strings.TrimPrefix(strings.TrimSuffix(match[1], "]]>"), "<![CDATA[")

			if !seenCreators[creator] {
				glogger.Done("Creator:", creator)
				creators = append(creators, creator)
				seenCreators[creator] = true
			}
		}
	}
	return creators
}
