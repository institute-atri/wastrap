/*
Package update provides functionality for updating the application.
*/
package update

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/institute-atri/ghttp"
	"github.com/institute-atri/glogger"
	"gopkg.in/yaml.v3"
)

// The Application represents the structure of the application version.
type Application struct {
	Application struct {
		Version string `yaml:"version"`
	} `yaml:"application"`
}

func CheckRepositoryVersion() string {
	url := "https://raw.githubusercontent.com/institute-atri/wastrap/main/internal/config/config.yaml"
	var response = ghttp.GET(url)

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

	if appInfos.Application.Version == "" {
		glogger.Fatal("Failed to find update, check github link: https://github.com/institute-atri/wastrap")
	}

	return appInfos.Application.Version
}

func FindConfigFile(test bool) string {
	var filePath string

	if test {
		filePath = "../config/config.yaml"
	} else {
		filePath = "internal/config/config.yaml"
	}

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

// CheckUpdate checks if an update is available by comparing the repository version with the current version.
// If an update is available, it triggers the update process.
func CheckUpdate() error {
	repositoryVersion := CheckRepositoryVersion()
	programVersion := FindConfigFile(false)

	if repositoryVersion != programVersion && programVersion != "" && repositoryVersion != "" {
		updateWastrapPermission := glogger.ScanQ("Do you want to update wastrap [Y/n] ")
		GettingUpdate(updateWastrapPermission)
		return nil
	}

	err := errors.New("failed to find version")
	return err
}
