package firewalls

import "testing"

func TestCheckWAFPlugins(t *testing.T) {
	cases := []struct {
		name    string
		url     string
		isError bool
	}{
		{"EmptyURL", "", true},
		{"InvalidURL", "not_a_valid_url", true},
		{"ValidURL", "http://www.santosfc.com.br/", true},
	}

	for _, result := range cases {
		t.Run(result.name, func(t *testing.T) {
			CheckWAFPlugins(result.url)

			if !result.isError {
				t.Errorf("expected no error but got logs: %v", result.url)
			}
		})
	}
}
