package config

import "testing"

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
