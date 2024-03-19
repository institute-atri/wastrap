package users

import (
	"regexp"
	"strings"

	"github.com/institute-atri/glogger"
	"github.com/institute-atri/gnet"
)

func Passive(url string) {
	urlFeed := url + "/feed"

	response := gnet.GET(urlFeed)

	search := regexp.MustCompile(`<dc:creator><!\[CDATA\[(.+?)\]\]></dc:creator>`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	if len(matches) == 0 {
		glogger.Danger("Nothing found")
	} else {
		seenCreators := make(map[string]bool)
		for _, match := range matches {
			creators := strings.TrimPrefix(strings.TrimSuffix(match[1], "]]>"), "<![CDATA[")
	
			if !seenCreators[creators] {
				glogger.Done("Creator:", creators)
				seenCreators[creators] = true
			}
		}
	}
}
