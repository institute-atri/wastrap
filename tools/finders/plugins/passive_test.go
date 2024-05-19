package plugins

import "testing"

func TestPassive(t *testing.T) {
	Passive("https://test-atri.netlify.app/")
}

func TestExtractPluginNameTrue(t *testing.T){
	response := extractPluginName("https://test-atri.netlify.app/wp-content/plugins/ml-slider/assets/metaslider/script.min.js")
	if response != "ml-slider"{
		t.Fatal("What was expected was ml-slider but it was:", response)
	}
}

func TestExtractPluginNameFalse(t *testing.T){
	response := extractPluginName("https://test-atri.netlify.app/wp-content")
	if response != ""{
		t.Fatal("What was expected was '' but it was:", response)
	}
}