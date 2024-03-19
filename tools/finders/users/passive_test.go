package users

import "testing"

func TestPassive(t *testing.T) {
	creators := Passive("https://tests-wp-atri.vercel.app")

	if len(creators) != 1 {
		t.Errorf("The number of users is: %d, expected was 1", len(creators))
	}
}