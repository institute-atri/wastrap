package validator

import (
	"regexp"
	"strings"
	"github.com/institute-atri/glogger"
	"github.com/institute-atri/ghttp"
)

// It must perform validation to determine if the target inserted through the --url parameter of the Command Line Interface (CLI) is effectively running the WordPress framework.
func Wordpress(url string) bool {
	var confidence float32

	var payloads = [4]string{
		"<meta name=\"generator content=\"WordPress",
		"<a href=\"http://www.wordpress.com\">Powered by WordPress</a>",
		"<link rel=\"https://api.wordpress.org/",
		"<link rel=\"https://api.w.org/\"",
	}

	var response = ghttp.GET(url)

	for _, payload := range payloads {
		if strings.Contains(response.BRaw, payload) {
			confidence++
		}
	}
	return confidence/4*100 <= 50
}
func WordpressVersion(url string)string{
	request := ghttp.NewHttp()

	request.SetURL(url)
	request.SetMethod("GET")

	response, err := request.Do()

	if err != nil {
		glogger.Danger(err)
	}

	rex := regexp.MustCompile("<meta name=\"generator\" content=\"WordPress ([0-9.-]*).*?")
	return rex.FindStringSubmatch(response.BRaw)[1]
}
