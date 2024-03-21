package update

import (
	"strings"
	"testing"
)

// hasTheLetterV checks if a given string contains the letter 'v'.
// It iterates over each character in the input string and returns true if 'v' is found, otherwise returns false.
func hasTheLetterV(versionExpectation string) bool {
	for _, char := range versionExpectation {
		if char == 'v' {
			return true
		}
	}
	return false
}

// TestCheckRepositoryVersion verifies the behavior of the CheckRepositoryVersion function, which retrieves the version of a repository from a specified URL.
// It calls the CheckRepositoryVersion function and checks if the returned version string contains the letter 'v'.
// This test ensures that the CheckRepositoryVersion function retrieves the repository version and returns a string containing the letter 'v', as expected.
func TestCheckRepositoryVersion(t *testing.T) {
	checkRepositoryVersionReturn := CheckRepositoryVersion()
	expected := hasTheLetterV(checkRepositoryVersionReturn)

	if !expected {
		t.Errorf("expected %v to be present in CheckRespositoryVersion, but it is returning %s", expected, checkRepositoryVersionReturn)
	}
}

// FindConfigFile checks if a config file exists based on the provided boolean parameter.
// If the parameter is true, it attempts to find the config file; otherwise, it returns an empty string.
func TestFindConfigFile(t *testing.T) {
	findConfigFileReturn := FindConfigFile(true)
	expected := hasTheLetterV(findConfigFileReturn)

	if !expected {
		t.Errorf("expected %v to be present in findConfigFileReturn, but it is returning %s", expected, findConfigFileReturn)
	}
}

// TestCheckUpdate checks the functionality of the CheckUpdate function.
// It verifies that the function returns an error containing the message "failed to find version" when the update check fails.
func TestCheckUpdate(t *testing.T) {
	err := CheckUpdate()
	if err != nil && !strings.Contains(err.Error(), "failed to find version") {
		t.Errorf("The error was: %v", err)
	}
}
