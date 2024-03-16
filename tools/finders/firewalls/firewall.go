package firewalls

import (
	"strconv"
	"strings"

	"github.com/institute-atri/glogger"
	"github.com/institute-atri/gnet"
)

var wafPlugins = []string{
	"wordfence",
	"wordfence-assistant",
	"cloudflare",
	"bulletproof-security",
	"better-wp-security",
	"sucuri-scanner",
	"wp-security-scan",
	"block-bad-queries",
	"all-in-one-wp-security-and-firewall",
	"6scan-protection",
	"siteguard",
	"ninjafirewall",
	"malcare-security",
	"wp-cerber",
	"wesecur-security",
	"itheme-security",
	"vaultpress",
	"flaming-firewall",
	"sucuri-security",
	"wp-firewall",
	"Shield-Security",
	"Security-Ninja",
	"iThemes-Security",
	"Sucuri",
	"MalCare",
	"JetPack",
}

func CheckWAFPlugins(siteURL string) {
	if siteURL == "" {
		glogger.Danger("Please provide a URL")
		return
	}
	if !strings.HasPrefix(siteURL, "http://") && !strings.HasPrefix(siteURL, "https://") {
		glogger.Danger("The link provided is not a valid URL\n")
		return
	}

	pluginsURL := siteURL + "/wp-content/plugins/"
	glogger.Info("Detected Web Application Firewalls (WAFs):\n")

	for _, plugin := range wafPlugins {
		url := pluginsURL + plugin + "/"

		response := gnet.GET(url)
		responseStatus := strconv.Itoa(response.StatusCode)
		messageResponse := "Plugin: \"" + plugin + "\" | Status: \"" + responseStatus + "\""

		if response.StatusCode == 200 {
			glogger.Done(messageResponse)
		} else {
			glogger.Warning(messageResponse)
		}
	}
}
