package validator

import (
	"strings"

	"github.com/institute-atri/ghttp"
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
