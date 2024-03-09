package update

import (
	"io"
	"os"
	"strings"

	"github.com/institute-atri/glogger"
	"github.com/institute-atri/gnet"
	"gopkg.in/yaml.v3"
)

type Application struct {
	Application struct {
		Version string `yaml:"version"`
	} `yaml:"application"`
}

func checkRepositoryVersion() string {
	url := "https://raw.githubusercontent.com/institute-atri/wastrap/main/internal/config/config.yaml"
	var response = gnet.GET(url)

	reader := strings.NewReader(response.BRaw)
	body, err := io.ReadAll(reader)
	if err != nil {
		glogger.Danger("Failed to find update, check github link: https://github.com/institute-atri/wastrap")
		return ""
	}

	var appInfos Application
	err = yaml.Unmarshal(body, &appInfos)
	if err != nil {
		glogger.Danger("Failed to find update, check github link: https://github.com/institute-atri/wastrap")
		return ""
	}

	return appInfos.Application.Version
}

func findConfigFile() string {
	filePath := "internal/config/config.yaml"

	file, err := os.ReadFile(filePath)
	if err != nil {
		glogger.Danger("The program is damaged, check github link: https://github.com/institute-atri/wastrap")
		return ""
	}

	var appInfo Application
	err = yaml.Unmarshal(file, &appInfo)
	if err != nil {
		glogger.Danger("The program is damaged, check github link: https://github.com/institute-atri/wastrap")
		return ""
	}

	return appInfo.Application.Version
}

func CheckUpdate() {
	var respositoryVersion string = checkRepositoryVersion()
	var programVersion string = findConfigFile()

	if respositoryVersion != programVersion && programVersion != "" && respositoryVersion != "" {
		GettingUpdate()
	}

}
