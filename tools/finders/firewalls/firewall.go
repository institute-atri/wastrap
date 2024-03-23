package firewalls

import (
	"strconv"

	"github.com/institute-atri/ghttp"
	"github.com/institute-atri/glogger"
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
	pluginsURL := siteURL + "/wp-content/plugins/"
	glogger.Info("Detected Web Application Firewalls (WAFs):\n")

	for _, plugin := range wafPlugins {
		url := pluginsURL + plugin + "/"

		response := ghttp.GET(url)
		responseStatus := strconv.Itoa(response.StatusCode)
		messageResponse := "Plugin: \"" + plugin + "\" | Status: \"" + responseStatus + "\""

		if response.StatusCode == 200 {
			glogger.Done(messageResponse)
		} else {
			glogger.Warning(messageResponse)
		}
	}
}
