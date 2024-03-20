package users

import "testing"

func TestAgressive(t *testing.T) {
	result := Agressive("https://tests-wp-atri.vercel.app")
	if len(result) != 3 {
		t.Errorf("The number of result is: %d, expected was 3", len(result))
	}
}

func TestMethodAuthorSitemapXML(t *testing.T) {
	authors := methodAuthorSitemapXML("https://tests-wp-atri.vercel.app/local2")
	if len(authors) != 2 {
		t.Errorf("The number of authors is: %d, expected was 2", len(authors))
	}

	returnInvalid := methodAuthorSitemapXML("https://tests-wp-atri.vercel.app/invalid")
	if len(returnInvalid) != 0 {
		t.Errorf("The number of authors is: %d, expected was 0", len(returnInvalid))
	}
}

func TestMethodWpJson(t *testing.T) {
	slugs := methodWpJson("https://tests-wp-atri.vercel.app")
	if len(slugs) != 2 {
		t.Errorf("The number of slugs is: %d, expected was 2", len(slugs))
	}

	returnInvalid := methodWpJson("https://tests-wp-atri.vercel.app/invalid")
	if len(returnInvalid) != 0 {
		t.Errorf("The number of slugs is: %d, expected was 0", len(returnInvalid))
	}
}

func TestMethodRestRouteWp(t *testing.T) {
	returnInvalid := methodRestRouteWp("https://tests-wp-atri.vercel.app")
	if len(returnInvalid) != 0 {
		t.Errorf("The number of slugs is: %d, expected was 0", len(returnInvalid))
	}
}