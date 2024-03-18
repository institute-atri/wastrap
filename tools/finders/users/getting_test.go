package users

import "testing"

func TestGetting(t *testing.T) {
	users := Getting("https://tests-wp-atri.vercel.app")
	if len(users) != 2 {
		t.Errorf("The number of users is: %d, expected was 2", len(users))
	}
}
