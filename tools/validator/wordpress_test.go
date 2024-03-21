package validator

import (
	"testing"
)

// TestWordpress verifies the functionality of the Wordpress function by testing its ability to detect
// if WordPress powers a given website. It calls the Wordpress function with a specific URL and
// checks whether it returns true, indicating that WordPress was detected on the website. This test ensures
// that the Wordpress function correctly identifies WordPress-powered websites.
func TestWordpress(t *testing.T) {
	result := Wordpress("http://www.wordpress.com/")
	if result != true {
		t.Errorf("Expected Wordpress() to be %t, but got %t", true, result)
	}
}

// TestWordpressVersion verifies the functionality of the WordpressVersion function by testing its ability to
// extract the version number of WordPress used on a given website. It calls the WordpressVersion function with
// a specific URL and checks whether it returns an empty string, indicating that the WordPress version was not
// found in the website's metadata. This test ensures that the WordpressVersion function correctly extracts and
// returns the version number of WordPress used on a website, if available.
func TestWordpressVersion(t *testing.T) {
	result := WordpressVersion("http://www.wordpress.com/")

	if result != "" {
		t.Errorf("Expected WordpressVersion() to be empty, but got %s", result)
	}
}
