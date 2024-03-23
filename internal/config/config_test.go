package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetVersion(t *testing.T) {

	cfg := Config{
		Application: struct {
			Name        string `yml:"name"`
			Version     string `yml:"version"`
			Author      string `yml:"author"`
			Description string `yml:"description"`
		}{
			Version: "1.0.0",
		},
	}

	//call the function GetVersion
	version := GetVersion(cfg)

	expectedVersion := "1.0.0"
	if version != expectedVersion {
		t.Errorf("Incorrect version. Expeted: %s, Obtained: %s", expectedVersion, version)
	}
}

func TestLoadConfig(t *testing.T) {
	tempFile := "test_config.yml"
	err := ioutil.WriteFile(tempFile, []byte(`
application:
  name: Test App
  version: v1.0.0
  author: Test Author
  description: Test Description
`), 0644)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(tempFile)

	cfg, err := LoadConfig(tempFile)
	if err != nil {
		t.Fatalf("Error loading configuration: %v", err)
	}

	if cfg.Application.Name != "Test App" {
		t.Errorf("Incorrect application name. Expected: %s, Obtained: %s", "Test App", cfg.Application.Name)
	}
	if cfg.Application.Version != "v1.0.0" {
		t.Errorf("Incorrect application version. Expeted: %s, Obtained: %s", "v1.0.0", cfg.Application.Version)
	}

}
