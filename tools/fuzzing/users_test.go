package fuzzing

import "testing"

// TestPassive verifies the functionality of the Passive function by simulating a passive brute force attack
// on a WordPress URL provided, using a specific username and a list of passwords. This test focuses on
// ensuring the correct behavior of the Passive function without relying on actual network connections or mocks.
func TestPassive(t *testing.T) {
	url := "http://www.wordpress.com/"
	username := "testuser"
	passwords := []string{"password1", "password2", "password3"}
	Passive(url, username, passwords)
}

// TestAggressive verifies if the Aggressive function works correctly by performing an aggressive brute force attack
// on a specific WordPress URL, with a provided username and a list of passwords. This test does not make use of mocks
// or real HTTP connections, as it focuses solely on the internal logic of the Aggressive function.
func TestAggressive(t *testing.T) {
	url := "http://www.wordpress.com/"
	username := "testuser"
	passwords := []string{"password1", "password2", "password3"}

	Aggressive(url, username, passwords)
}
