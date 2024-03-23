package validator

import (
	"regexp"
	"strings"

	"github.com/institute-atri/ghttp"
	"github.com/institute-atri/glogger"
)

func Wordpress(url string) bool {
	var confidence float32

	var payloads = [4]string{
		"<meta name=\"generator content=\"WordPress",
		"<a href=\"http://www.wordpress.com\">Powered by WordPress</a>",
		"<link rel=\"https://api.wordpress.org/",
		"<link rel=\"https://api.w.org/\"",
	}

	response := ghttp.GET(url)

	for _, payload := range payloads {
		if strings.Contains(response.BRaw, payload) {
			confidence++
		}
	}
	return confidence/4*100 <= 50
}
func WordpressVersion(url string) string {
	request := ghttp.NewHttp()

	request.SetURL(url)
	request.SetMethod("GET")

	response, err := request.Do()

	if err != nil {
		glogger.Danger(err)
	}

	rex := regexp.MustCompile("<meta name=\"generator\" content=\"WordPress ([0-9.-]*).*?")
	match := rex.FindStringSubmatch(response.BRaw)
	if len(match) < 2 {
		glogger.Warning("Wordpress version not found")
		return ""
	}
	return match[1]
}
