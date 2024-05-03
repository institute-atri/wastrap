package plugins

import (
	"regexp"

	"github.com/institute-atri/ghttp"
	"github.com/institute-atri/glogger"
)

var pluginsWithReadme = make(map[string]bool)

func status_link(url string) *ghttp.HttpResponse {
	url += "/wp-content/plugins"
	response := ghttp.GET(url)
	return response
}

func search_readme(url string, plugin string) bool {
	url += "/wp-content/plugins/" + plugin
	response := ghttp.GET(url)

	search := regexp.MustCompile(`[^'"]*?readme\.txt`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	for range matches {
		pluginsWithReadme[plugin] = true
		return true
	}
	pluginsWithReadme[plugin] = false
	return false
}

func collect_readme_version(url string, plugin string) string {
	url += "/wp-content/plugins/" + plugin + "/readme.txt"

	response := ghttp.GET(url)

	search := regexp.MustCompile(`Stable tag: (\d+\.\d+\.\d+)`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	for _, match := range matches {
		return match[1]
	}
	return ""
}


func search_changelog(url string, plugin string) bool {
	url += "/wp-content/plugins/" + plugin
	response := ghttp.GET(url)

	search := regexp.MustCompile(`[^'"]*?changelog\.txt`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	for range matches {
		return true
	}
	return false
}

func collect_changelog_version(url string, plugin string) string {
	url += "/wp-content/plugins/" + plugin + "/changelog.txt"

	response := ghttp.GET(url)

	search := regexp.MustCompile(`(\d+\.\d+\.\d+)`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	for _, match := range matches {
		return match[0]
	}
	return ""
}

func Agressive(url string) {
	response := status_link(url)

	if response.StatusCode != 200 {
		glogger.Danger("Unable to access the page")
	}

	search := regexp.MustCompile(`href=["']([^"']+)["']`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	for _, match := range matches {
		search_readme(url, match[1])
	}

	glogger.Done("======Result======")

	for plugin, has_readme := range pluginsWithReadme {
		if has_readme {
			version := collect_readme_version(url, plugin)
			glogger.Done("- ", plugin)
			println("  | Version: ", version)
			println("     | Method: readme.txt")
		} else {
			changelog := search_changelog(url, plugin)
			if changelog {
				version := collect_changelog_version(url, plugin)
				glogger.Done("- ", plugin)
				println("  | Version: ", version)
				println("     | Method: changelog.txt")
			} 

		}
	}
}
