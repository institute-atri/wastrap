package update

import (
	"os"
	"path/filepath"
	"testing"
)

// TestCheckTheGitPath verifies the behavior of the checkTheGitPath function in determining the presence of a .git directory.
// It creates a temporary directory and adds a .git directory inside it to simulate the presence of a Git repository for positive testing.
// The function is then tested for both cases: when the .git directory exists and when it does not exist in the current directory.
// This test ensures that the checkTheGitPath function correctly identifies the presence or absence of a .git directory.
func TestCheckTheGitPath(t *testing.T) {
	tempDir := t.TempDir()

	gitDir := filepath.Join(tempDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatalf("error creating .git directory: %v", err)
	}

	if !checkTheGitPath(true) {
		t.Error("expected true, got false for existing .git directory in current directory")
	}
	if checkTheGitPath(false) {
		t.Error("expected false, got true for non-existing .git directory in parent directory")
	}
}

// TestCheckIfGitIsInstalled verifies the behavior of the checkIfGitIsInstalled function in determining whether Git is installed on the system.
// It checks if Git is installed by calling the checkIfGitIsInstalled function and expects a true result indicating Git is installed.
// This test ensures that the checkIfGitIsInstalled function correctly identifies whether Git is installed on the system.
func TestCheckIfGitIsInstalled(t *testing.T) {
	result := checkIfGitIsInstalled()
	if result == false {
		t.Errorf("Expected: %v but bot: %v", false, result)
	}
}

// TestOperationSystem verifies the behavior of the operating system detection functions: isMacOS, isWindows, isDebian, and isFedora.
// It checks each operating system detection function individually and compares the result with the expected value.
// This test ensures that the operating system detection functions correctly identify the current operating system.
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

// TestPermissionIf verifies the behavior of the permissionIf function, which checks if a given input represents affirmative permission (e.g., "yes", "y", "YES", "Y") or negative permission (e.g., "no", "n", "NO", "N").
// It tests the permissionIf function with various input cases, including affirmative and negative permission strings, and compares the returned value with the expected result.
// This test ensures that the permissionIf function correctly identifies affirmative and negative permissions based on the input string.
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

// TestInstallGitOS verifies the installation of Git on different operating systems.
// It calls the installation functions for macOS, Debian, and Fedora to ensure that Git installation procedures for these operating systems do not return errors.
// This test ensures that Git can be installed successfully on various supported operating systems.
func TestInstallGitOS(t *testing.T) {
	installGitMacOs()
	installGitDebian()
	installGitFedora()
}

// TestInstallGitWindows verifies the behavior of the installGitWindows function in a Windows environment.
// It temporarily modifies the PATH environment variable to an empty value to prevent the system from attempting to execute unknown commands.
// Then, it tests the function with different user inputs to ensure correct behavior.
// Finally, it restores the original value of the PATH environment variable after the test to maintain system integrity.
func TestInstallGitWindows(t *testing.T) {
	originalPath := os.Getenv("PATH")
	os.Setenv("PATH", "")

	defer func() {
		os.Setenv("PATH", originalPath)
	}()

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

// TestInstallGitUnknown verifies the behavior of the installGitUnknown function when the operating system is unknown.
// It temporarily modifies the PATH environment variable to an empty value to prevent the system from attempting to execute unknown commands.
// Then, it tests the function with different user inputs to ensure correct behavior.
// Finally, it restores the original value of the PATH environment variable after the test to maintain system integrity.
func TestInstallGitUnknown(t *testing.T) {
	originalPath := os.Getenv("PATH")
	os.Setenv("PATH", "")

	defer func() {
		os.Setenv("PATH", originalPath)
	}()

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

// TestUpdateSoftware verifies the behavior of the updateSoftware function when updateWithGit is set to true.
// It ensures that the updateWithGit function is not called and createGitPath function is not called.
// This test aims to confirm the correct behavior of updateSoftware under specific conditions.
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

// TestGettingUpdate verifies the behavior of the GettingUpdate function with different user permissions.
// It runs subtests for each permission value (Y/n) to simulate user input and test the corresponding behavior of the GettingUpdate function.
// This test ensures that the GettingUpdate function behaves correctly based on the provided permissions.
func TestGettingUpdate(t *testing.T) {
	tests := []struct {
		permission string
	}{
		{"Y"},
		{"n"},
	}

	for _, tt := range tests {
		t.Run(tt.permission, func(t *testing.T) {
			GettingUpdate(tt.permission)
		})
	}
}

// TestUpdateWithGit verifies the behavior of the updateWithGit function under different conditions.
// It tests the function with both true and false parameters to check if updates are attempted when enabled and if errors are returned when disabled.
// This test ensures that the updateWithGit function behaves correctly in updating software with Git based on the provided parameter.
func TestUpdateWithGit(t *testing.T) {
	if err := updateWithGit(true); err != nil {
		t.Errorf("updateWithGit(true) failed, expected no error, got: %v", err)
	}

	if err := updateWithGit(false); err != nil {
		t.Errorf("updateWithGit(false) failed, expected error, got nil")
	}
}
