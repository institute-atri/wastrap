package finders

import (
	"net/http"
	"strconv"
	"sync"
	"time"

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

	var wg sync.WaitGroup
	wg.Add(len(wafPlugins))

	glogger.Println("Detected Web Application Firewalls (WAFs):")

	for _, plugin := range wafPlugins {
		go func(plugin string) {
			defer wg.Done()

			url := pluginsURL + plugin + "/"

			startTime := time.Now()

			resp, err := http.Head(url)
			if err != nil {
				errorMessage := "Error accessing the plugin: \"" + plugin + "\", " + err.Error()
				glogger.Danger(errorMessage)
				return
			}

			endTime := time.Now()
			elapsedTime := endTime.Sub(startTime)
			glogger.Info("Time taken to access plugin \"" + plugin + "\": " + elapsedTime.String())

			statusMessage := http.StatusText(resp.StatusCode)
			message := "Response status for plugin: \"" + plugin + "\" = " + strconv.Itoa(resp.StatusCode) + " " + statusMessage + "\n"

			if resp.StatusCode == 200 {
				glogger.Done(message)
			} else {
				glogger.Warning(message)
			}

		}(plugin)
	}

	wg.Wait()
}
