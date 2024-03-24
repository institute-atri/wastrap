package plugins

import (
	"fmt"
	"regexp"

	"github.com/institute-atri/ghttp"
	"github.com/institute-atri/glogger"
)

func Passive(url string) {
	response := ghttp.GET(url)

	search := regexp.MustCompile(`<script[^>]*src="([^"]*plugin[^"]+)\?ver=([^"]+)"[^>]*></script>`)

	matches := search.FindAllStringSubmatch(response.BRaw, -1)

	equalVersions := make(map[string]bool)
	pluginVersions := make(map[string]map[string][]string)
	pluginTotalLinks := make(map[string]int)

	for _, match := range matches {
		if !equalVersions[match[2]] {
			equalVersions[match[2]] = false

			link := match[1]
			version := match[2]
			pluginName := extractPluginName(link)

			if pluginVersions[pluginName] == nil {
				pluginVersions[pluginName] = make(map[string][]string)
				pluginTotalLinks[pluginName] = 0
			}

			if pluginVersions[pluginName][version] == nil {
				pluginVersions[pluginName][version] = []string{}
			}

			pluginVersions[pluginName][version] = append(pluginVersions[pluginName][version], link)
			pluginTotalLinks[pluginName]++
		}
	}

	for pluginName, versions := range pluginVersions {
		if len(versions) > 1 {
			glogger.Done(pluginName)
			for version, links := range versions {
				percentage := float64(len(links)) / float64(pluginTotalLinks[pluginName]) * 100
				println("| Version " + version + ":")
				println("  | Confidence: " + fmt.Sprintf("%.2f", percentage) + "%")
				println("  | Match:")
				for _, link := range links {
					println("    |" + link)
				}
			}
		} else {
			glogger.Done(pluginName)
			for version, links := range versions {
				percentage := float64(len(links)) / float64(pluginTotalLinks[pluginName]) * 100
				println("| Version " + version + ":")
				println("  | Confidence: " + fmt.Sprintf("%.2f", percentage) + "%")
				println("  | Match:")
				for _, link := range links {
					println("    | - " + link)
				}
			}
		}
	}
}

func extractPluginName(link string) string {
	search := regexp.MustCompile(`plugins/([^/]+)/`)
	matches := search.FindStringSubmatch(link)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
