package update

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Application struct {
	Application struct {
		Version string `yaml:"version"`
	} `yaml:"application"`
}

func checkRepositoryVersion() string {
	url := "https://raw.githubusercontent.com/institute-atri/wastrap/main/internal/config/config.yaml"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to find update, check github link: https://github.com/institute-atri/wastrap")
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to find update, check github link: https://github.com/institute-atri/wastrap")
		return ""
	}

	var appInfos Application
	err = yaml.Unmarshal(body, &appInfos)
	if err != nil {
		fmt.Println("Failed to find update, check github link: https://github.com/institute-atri/wastrap")
		return ""
	}

	return appInfos.Application.Version
}

func findConfigFile() string {
	filePath := "internal/config/config.yaml"

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("The program is damaged,: https://github.com/institute-atri/wastrap")
		return ""
	}

	var appInfo Application
	err = yaml.Unmarshal(file, &appInfo)
	if err != nil {
		fmt.Println("The program is damaged,: https://github.com/institute-atri/wastrap")
		return ""
	}

	return appInfo.Application.Version
}

func CheckUpdate() {
	var respositoryVersion string = checkRepositoryVersion()
	var programVersion string = findConfigFile()

	if respositoryVersion != programVersion {
		GettingUpdate()
	}

}
