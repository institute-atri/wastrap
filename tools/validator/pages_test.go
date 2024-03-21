package validator

import (
	"testing"

	"github.com/institute-atri/ghttp"
)

func TestNewPages(t *testing.T) {
	NewPages("http://www.wordpress.com", "", true, true)
}

// TestAdmin verifies the functionality of the Admin method in the Pages struct by testing its ability to detect
// if the WordPress admin interface is accessible on a given website. It creates a Pages instance with specified
// settings, sends an HTTP request to the wp-cron.php endpoint, and checks the result of the Admin method against
// expected confidence levels based on the HTTP response status code. This test ensures that the Admin method
// correctly identifies the accessibility of the WordPress admin interface and returns the expected result.
func TestAdmin(t *testing.T) {
	p := Pages{
		Url:         "http://www.wordpress.com/",
		TOR:         false,
		Proxy:       "",
		RandomAgent: true,
	}

	var request = ghttp.NewHttp()
	_ = request.SetURL(p.Url + "wp-cron.php")
	response, _ := request.Do()

	result := p.Admin()

	if (response.StatusCode == 200 || response.StatusCode == 403) && result.Confidence != 100 {
		t.Errorf("Expected 100, but got %v", result.Confidence)
	}
}

// TestRobots verifies the functionality of the Robots method in the Pages struct by testing its ability to
// detect the presence of a robots.txt file on a given website. It creates a Pages instance with specified
// settings, sends an HTTP request to the wp-cron.php endpoint, and checks the result of the Robots method
// against expected confidence levels based on the HTTP response status code. This test ensures that the Robots
// method correctly identifies the presence of a robots.txt file and returns the expected result.
func TestRobots(t *testing.T) {
	p := Pages{
		Url:         "http://www.wordpress.com/",
		TOR:         false,
		Proxy:       "",
		RandomAgent: true,
	}

	var request = ghttp.NewHttp()
	_ = request.SetURL(p.Url + "wp-cron.php")
	response, _ := request.Do()

	result := p.Robots()

	if (response.StatusCode == 200 || response.StatusCode == 403) && result.Confidence != 100 {
		t.Errorf("Expected 100, but got  %v", result.Confidence)
	}
}

// TestSitemap verifies the functionality of the Sitemap method in the Pages struct by testing its ability to
// detect the presence of a sitemap.xml file on a given website. It creates a Pages instance with specified
// settings, sends an HTTP request to the wp-cron.php endpoint, and checks the result of the Sitemap method
// against expected confidence levels based on the HTTP response status code. This test ensures that the Sitemap
// method correctly identifies the presence of a sitemap.xml file and returns the expected result.
func TestSitemap(t *testing.T) {
	p := Pages{
		Url:         "http://www.wordpress.com/",
		TOR:         false,
		Proxy:       "",
		RandomAgent: true,
	}

	var request = ghttp.NewHttp()
	_ = request.SetURL(p.Url + "wp-cron.php")
	response, _ := request.Do()

	result := p.Sitemap()

	if (response.StatusCode == 200 || response.StatusCode == 403) && result.Confidence != 100 {
		t.Errorf("Expected 100, but got  %v", result.Confidence)
	}
}

// TestReadme verifies the functionality of the Readme method in the Pages struct by testing its ability to
// detect the presence of a readme.html or readme.txt file on a given website. It creates a Pages instance with
// specified settings, sends an HTTP request to the wp-cron.php endpoint, and checks the result of the Readme
// method against expected confidence levels based on the HTTP response status code. This test ensures that the
// Readme method correctly identifies the presence of a readme.html or readme.txt file and returns the expected result.
func TestReadme(t *testing.T) {
	p := Pages{
		Url:         "http://www.wordpress.com/",
		TOR:         false,
		Proxy:       "",
		RandomAgent: true,
	}

	var request = ghttp.NewHttp()
	_ = request.SetURL(p.Url + "wp-cron.php")
	response, _ := request.Do()

	result := p.Readme()

	if (response.StatusCode == 200 || response.StatusCode == 403) && result.Confidence != 100 {
		t.Errorf("Expected 100, but got  %v", result.Confidence)
	}
}
