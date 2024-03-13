package update

import (
	"os"
	"path/filepath"
	"runtime"
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
	CheckRepositoryVersionReturn := CheckRepositoryVersion()
	expected := hasTheLetterV(CheckRepositoryVersionReturn)

	if !expected {
		t.Errorf("expected 'vx.x.x' to be present in CheckRespositoryVersion, but it is returning %s", CheckRepositoryVersionReturn)
	}
}

func TestFindConfigFile(t *testing.T) {
	FindConfigFileReturn := FindConfigFile(true)
	expected := hasTheLetterV(FindConfigFileReturn)

	if !expected {
		t.Errorf("expected 'vx.x.x' to be present in findConfigFileReturn, but it is returning %s", FindConfigFileReturn)
	}
}

func TestCheckUpdate(t *testing.T) {
	err := CheckUpdate()
	if err != nil && !strings.Contains(err.Error(), "failed to find version"){
		t.Errorf("The error was: %v", err)
	}
}

func TestCheckTheGitPath(t *testing.T) {
	tmpDir := t.TempDir()
	currentDir, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(currentDir)

	got := checkTheGitPath()
	if got {
		t.Errorf("checkTheGitPath() = %v, want false", got)
	}

	gitDir := filepath.Join(tmpDir, ".git")
	os.Mkdir(gitDir, os.ModePerm)

	got = checkTheGitPath()
	if !got {
		t.Errorf("checkTheGitPath() = %v, want true", got)
	}
}

func TestCheckIfGitIsInstalled(t *testing.T) {
	if !checkIfGitIsInstalled() {
		t.Error("Expected true, got false")
	}
}

func TestOperationgSystem(t *testing.T) {
	type Cases struct {
		Variable string
		Expected bool
	}

	var cases []Cases = []Cases{
		{"darwin", false},
		{"windows", false},
		{"debian", false},
		{"fedora", false},
	}

	if runtime.GOOS == "linux" {
		_, errDebian := os.Stat("/etc/debian_version")
		if errDebian == nil {
			cases[2].Expected = true
		} else {
			cases[3].Expected = true
		}
	} else if runtime.GOOS == "darwin" {
		cases[0].Expected = true
	} else if runtime.GOOS == "windows" {
		cases[1].Expected = true
	}

	if isMacOS() != cases[0].Expected {
		t.Errorf("isMacOs() = %t, want %t", isMacOS(), cases[0].Expected)
	}
	if isWindows() != cases[1].Expected {
		t.Errorf("isWindows() = %t, want %t", isWindows(), cases[1].Expected)
	}
	if isDebian() != cases[2].Expected {
		t.Errorf("isDebian() = %t, want %t", isDebian(), cases[2].Expected)
	}
	if isFedora() != cases[3].Expected {
		t.Errorf("isFedora() = %t, want %t", isDebian(), cases[3].Expected)
	}
}

func TestPermissionIf(t *testing.T) {
	cases := []struct {
		Variable string
		Expected bool
	}{
		{"y", true},
		{"Y", true},
		{"yes", true},
		{"YES", true},
		{"Yes", true},
		{"n", false},
		{"N", false},
		{"no", false},
		{"NO", false},
		{"No", false},
	}

	for _, c := range cases {
		if permissionIf(c.Variable) != c.Expected {
			t.Errorf("permissionIf(%q) = %v, want %v", c.Variable, !c.Expected, c.Expected)
		}
	}
}

func TestInstallingGit(t *testing.T) {
	returnFunc := installingGit()
	if returnFunc != nil {
		t.Error("no operating system was found")
	}
}

func TestInstallGitOS(t *testing.T) {
	installGitMacOs()
	installGitDebian()
	installGitFedora()
}

func TestInstallGitWindows(t *testing.T) {
	cases := []struct {
		Variable string
		Expected bool
	}{
		{"y", true},
		{"n", false},
	}

	for _, c := range cases {
		err := installGitWindows(c.Variable)
		if err != nil {
			t.Errorf("installGitWindows(%q) = %t, want %t", c.Variable, !c.Expected, c.Expected)
		}
	}
}

func TestInstallGitUnknow(t *testing.T) {
	cases := []struct {
		Variable string
		Expected bool
	}{
		{"y", true},
		{"n", false},
	}

	for _, c := range cases {
		err := installGitUnknown(c.Variable)
		if err != nil {
			t.Errorf("installGitUnknown(%q) = %v, want %v", c.Variable, !c.Expected, c.Expected)
		}
	}
}

func TestUpdateWithGit(t *testing.T) {
	err := updateWithGit(true)
	if err != nil {
		t.Errorf("An error occurred in the updateWithGit functio")
	}
}

func TestUpdateGitNotInstalled(t *testing.T) {
	updateGitNotInstalled()
}
