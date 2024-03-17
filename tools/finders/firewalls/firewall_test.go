package firewalls

import "testing"

func TestCheckWAFPlugins(t *testing.T) {
	CheckWAFPlugins("http://www.example.com/")
}
