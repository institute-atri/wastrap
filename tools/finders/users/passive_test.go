package users

import "testing"

func TestPassive(t *testing.T) {
	creators := Passive("https://tests-wp-atri.vercel.app")

	if len(creators) != 1 {
		t.Errorf("The number of creators is: %d, expected was 1", len(creators))
	}

	returnInvalid := Passive("https://tests-wp-atri.vercel.app/invalid")

	if len(returnInvalid) != 0 {
		t.Errorf("The number of creators is: %d, expected was 0", len(returnInvalid))
	}
}