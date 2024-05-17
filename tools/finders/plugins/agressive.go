package plugins

import (
	"regexp"

	"github.com/institute-atri/ghttp"
	"github.com/institute-atri/glogger"
)

var pluginsWithReadme = make(map[string]bool)

func getResponse(url, page string) *ghttp.HttpResponse {
	url += page
	response := ghttp.GET(url)
	return response
}

// searching readme.txt
func searchReadme(url string, plugin string) bool {
	response := getResponse(url, "/wp-content/plugins/" + plugin)

	search := regexp.MustCompile(`[^'"]*?readme\.txt`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	for range matches {
		pluginsWithReadme[plugin] = true
		return true
	}
	pluginsWithReadme[plugin] = false
	return false
}

// colleting the version through readme.txt
func collectReadmeVersion(url string, plugin string) string {
	response := getResponse(url, "/wp-content/plugins/" + plugin + "/readme.txt")

	search := regexp.MustCompile(`Stable tag: (\d+\.\d+\.\d+)`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	for _, match := range matches {
		return match[1]
	}
	return ""
}

// searching changelog.txt
func searchChangelog(url string, plugin string) bool {
	response := getResponse(url, "/wp-content/plugins/" + plugin)

	search := regexp.MustCompile(`[^'"]*?changelog\.txt`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	for range matches {
		return true
	}
	return false
}

// colleting the version through changelog.txt
func collectChangelogVersion(url string, plugin string) string {
	response := getResponse(url, "/wp-content/plugins/" + plugin + "/changelog.txt")

	search := regexp.MustCompile(`(\d+\.\d+\.\d+)`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	for _, match := range matches {
		return match[0]
	}
	return ""
}


// mode agressive of searching for plugin versions
func Agressive(url string) {
	response := getResponse(url, "/wp-content/plugins")

	if response.StatusCode != 200 {
		glogger.Danger("Unable to access the page")
	}

	search := regexp.MustCompile(`href=["']([^"']+)["']`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	for _, match := range matches {
		searchReadme(url, match[1])
	}

	glogger.Done("======Result======")

	for plugin, has_readme := range pluginsWithReadme {
		if has_readme {
			version := collectReadmeVersion(url, plugin)
			glogger.Done("- ", plugin)
			println("  | Version: ", version)
			println("     | Method: readme.txt")
		} else {
			changelog := searchChangelog(url, plugin)
			if changelog {
				version := collectChangelogVersion(url, plugin)
				glogger.Done("- ", plugin)
				println("  | Version: ", version)
				println("     | Method: changelog.txt")
			} 

		}
	}
}
