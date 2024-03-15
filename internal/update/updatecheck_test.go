package update

import (
	"strings"
	"testing"
)

func hasTheLetterV(versionExpectation string) bool {
	for _, char := range versionExpectation {
		if char == 'v' {
			return true
		}
	}
	return false
}
func TestCheckRepositoryVersion(t *testing.T) {
	checkRepositoryVersionReturn := CheckRepositoryVersion()
	expected := hasTheLetterV(checkRepositoryVersionReturn)

	if !expected {
		t.Errorf("expected %v to be present in CheckRespositoryVersion, but it is returning %s", expected, checkRepositoryVersionReturn)
	}
}
func TestFindConfigFile(t *testing.T) {
	findConfigFileReturn := FindConfigFile(true)
	expected := hasTheLetterV(findConfigFileReturn)

	if !expected {
		t.Errorf("expected %v to be present in findConfigFileReturn, but it is returning %s", expected, findConfigFileReturn)
	}
}
func TestCheckUpdate(t *testing.T) {
	err := CheckUpdate()
	if err != nil && !strings.Contains(err.Error(), "failed to find version") {
		t.Errorf("The error was: %v", err)
	}
}
