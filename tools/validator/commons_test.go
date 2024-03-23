package validator

import (
	"testing"

	"github.com/institute-atri/ghttp"
)

// TestNewCommons verifies the functionality of the NewCommons constructor function by creating a new Commons instance
// with specified URL, proxy, TOR, and RandomAgent settings. It checks whether the created Commons instance matches
// the expected configuration. This test ensures the correct initialization of a Commons struct with the provided parameters.
func TestNewCommons(t *testing.T) {
	url := "http://www.wordpress.com/"
	proxy := "proxy.wordpress.com"

	commons := NewCommons(url, proxy, false, true)
	expected := &Commons{Url: url, Proxy: proxy, TOR: false, RandomAgent: true}

	if *commons != *expected {
		t.Errorf("Expected %v, but got %v", expected, commons)
	}
}

// TestPHPDisabled verifies the functionality of the PHPDisabled method in the Commons struct by testing
// its ability to detect if PHP is disabled on a given website. It creates a Commons instance with specified
// settings and checks the result of the PHPDisabled method against the expected outcome. This test ensures
// that the PHPDisabled method correctly identifies PHP disabled status and returns the expected result.
func TestPHPDisabled(t *testing.T) {
	c := Commons{
		Url:         "http://www.wordpress.com/",
		Proxy:       "",
		TOR:         false,
		RandomAgent: true,
	}
	result := c.PHPDisabled()

	expected := Interesting{
		Confidence: 100,
		FoundBy:    "Direct Access",
		BRaw:       "",
	}

	if result.FoundBy != expected.FoundBy {
		t.Errorf("Expected FoundBy %s, but got %s", expected.FoundBy, result.FoundBy)
	}

	if result.Confidence != expected.Confidence && result.BRaw != "" {
		t.Errorf("Expected Confidence %d, but got %d with BRaw: %s", expected.Confidence, result.Confidence, result.BRaw)
	}
}

// TestXMLRPC verifies the functionality of the XMLRPC method in the Commons struct by testing its ability to
// detect XML-RPC vulnerabilities on a given website. It creates a Commons instance with specified settings,
// sends an HTTP request to the XML-RPC endpoint, and checks the result of the XMLRPC method against expected
// confidence levels based on the HTTP response status code. This test ensures that the XMLRPC method correctly
// identifies XML-RPC vulnerabilities and returns the expected confidence level.
func TestXMLRPC(t *testing.T) {
	c := Commons{
		Url:         "http://www.wordpress.com/",
		TOR:         false,
		Proxy:       "",
		RandomAgent: true,
	}

	var request = ghttp.NewHttp()
	_ = request.SetURL(c.Url + "xmlrpc.php")
	response, _ := request.Do()

	result := c.XMLRPC()

	if result.Confidence < 80 {
		t.Errorf("Expected 80, but got %v", result.Confidence)
	}

	if result.Confidence < 60 {
		t.Errorf("Expected 60, but got %v", result.Confidence)
	}

	if (response.StatusCode == 200 || response.StatusCode == 403 || response.StatusCode == 405) && result.Confidence < 20 {
		t.Errorf("Expected 20, but got %v", result.Confidence)
	}
}

// TestWPCron verifies the functionality of the WPCron method in the Commons struct by testing its ability to
// detect whether the WordPress cron system is enabled on a given website. It creates a Commons instance with
// specified settings and checks the result of the WPCron method against expected confidence levels and other
// attributes. This test ensures that the WPCron method correctly identifies the status of the WordPress cron
// system and returns the expected result.
func TestWPCron(t *testing.T) {
	c := Commons{
		Url:         "http://www.wordpress.com/",
		TOR:         false,
		Proxy:       "",
		RandomAgent: true,
	}

	result := c.WPCron()

	if (result.Confidence != 100) && (result.Confidence != 0) {
		t.Errorf("Expected confidence level to be 100 when WordPress cron system is enabled, but got %v", result.Confidence)
	}

	if result.BRaw != "" {
		t.Errorf("Expected BRaw field to not be empty, but it is %v", result.BRaw)
	}

	if result.FoundBy != "Direct Access" {
		t.Errorf("Expected FoundBy field to be 'Direct Access', but got %v", result.FoundBy)
	}
}
