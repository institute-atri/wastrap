package plugins

import "testing"

var link string = "https://test-atri.netlify.app" 

func TestGetResponse(t *testing.T) {
	response := getResponse(link, "/wp-content/plugins")

	if response.StatusCode != 200 {
		t.Fatal("What was expected was 200 but it was:", response.StatusCode)
	}
}

func TestSearchReadme(t *testing.T) {
	response := searchReadme(link, "newsletter")

	if response != true {
		t.Fatal("What was expected was true but it was:", response)
	}

	responseFalse := searchReadme(link, "elementor-pro")

	if responseFalse != false {
		t.Fatal("What was expected was false but it was:", response)
	}

}

func TestCollectReadmeVersion(t *testing.T){
	response := collectReadmeVersion(link, "newsletter")

	if response != "8.3.2" {
		t.Fatal("What was expected was 8.3.2 but it was:", response)
	}
	
	responseFalse := collectReadmeVersion(link, "elementor-pro")

	if responseFalse != "" {
		t.Fatal("What was expected was '' but it was:", response)
	}

}

func TestSearchChangelog(t *testing.T){
	response := searchChangelog(link, "elementor-pro")
	
	if response != true {
		t.Fatal("What was expected was true but it was:", response)
	}

	responseFalse := searchChangelog(link, "not-exist")
	
	if responseFalse != false {
		t.Fatal("What was expected was false but it was:", response)
	}
}

func TestCollectChangelogVersion(t *testing.T){
	response := collectChangelogVersion(link, "elementor-pro")

	if response != "3.18.2"{
		t.Fatal("What was expected was 8.3.2 but it was:", response)
	}

	responseFalse := collectChangelogVersion(link, "not-exist")
	
	if responseFalse != "" {
		t.Fatal("What was expected was '' but it was:", response)
	}
}

func TestAgressive(t *testing.T){
	Agressive(link)
}