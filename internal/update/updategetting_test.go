package update

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckTheGitPath(t *testing.T) {

	// Defining a temporary directory for the tests
	tempDir := t.TempDir()

	// Creating a .git directory inside the temporary directory for the positive test
	gitDir := filepath.Join(tempDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatalf("error creating .git directory: %v", err)
	}

	// Testing the case where the .git directory exists in the current directory
	if !checkTheGitPath(true) {
		t.Error("expected true, got false for existing .git directory in current directory")
	}

	// Testing the case where the .git directory does not exist in the current directory
	if checkTheGitPath(false) {
		t.Error("expected false, got true for non-existing .git directory in parent directory")
	}
}
func TestCheckIfGitIsInstalled(t *testing.T) {
	result := checkIfGitIsInstalled()
	if result == false {
		t.Errorf("Expected: %v but bot: %v", false, result)
	}
}
func TestOperationSystem(t *testing.T) {
	type Cases struct {
		Variable string
		Expected bool
	}

	cases := []Cases{
		{"darwin", false},
		{"windows", false},
		{"debian", false},
		{"fedora", false},
	}

	if isMacOS() {
		cases[0].Expected = true
	} else if isWindows() {
		cases[1].Expected = true
	} else if isDebian() {
		cases[2].Expected = true
	} else if isFedora() {
		cases[3].Expected = true
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
		t.Errorf("isFedora() = %t, want %t", isFedora(), cases[3].Expected)
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
	_ = installGitMacOs()
	_ = installGitDebian()
	_ = installGitFedora()
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
func TestInstallGitUnknown(t *testing.T) {
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
func TestUpdateSoftware(t *testing.T) {
	var updateWithGitCalled, createGitPathCalled bool

	updateSoftware(true)

	if updateWithGitCalled {
		t.Errorf("Expected updateWithGit to be called")
	}
	if createGitPathCalled {
		t.Errorf("Did not expect createGitPath to be called")
	}
}
func TestGettingUpdate(t *testing.T) {
	tests := []struct {
		permission string
	}{
		{"y"},
		{"n"},
	}

	for _, tt := range tests {
		t.Run(tt.permission, func(t *testing.T) {
			GettingUpdate(tt.permission)
		})
	}
}
func TestUpdateWithGit(t *testing.T) {
	if err := updateWithGit(true); err != nil {
		t.Errorf("updateWithGit(true) failed, expected no error, got: %v", err)
	}

	if err := updateWithGit(false); err != nil {
		t.Errorf("updateWithGit(false) failed, expected error, got nil")
	}
}
func TestUpdateGitNotInstalled(t *testing.T) {
	updateGitNotInstalled()
}
