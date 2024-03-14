package update

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/institute-atri/glogger"
)

func checkTheGitPath(test bool) bool {
		currentDir, err := os.Getwd()
		if err != nil {
			glogger.Danger("O programa está danificado, verifique o link do GitHub: https://github.com/institute-atri/wastrap")
		}

		var gitDir string
		
		if test {
			gitDir = filepath.Join(currentDir,  "..", "..", ".git")
		} else {
			gitDir = filepath.Join(currentDir, "..", ".git")
		}
	
		_, err = os.Stat(gitDir)
		if err == nil {
			return true
		} else if os.IsNotExist(err) {
			return false
		}
		return false
}

func checkIfGitIsInstalled() bool {
	cmd := exec.Command("git", "--version")
	err := cmd.Run()
	return err == nil
}

func isMacOS() bool {
	return runtime.GOOS == "darwin"
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func isDebian() bool {
	_, err := os.Stat("/etc/debian_version")
	return err == nil
}

func isFedora() bool {
	_, err := os.Stat("/etc/fedora-release")
	return err == nil
}

func openBrowser() {
	url := "https://git-scm.com/downloads"

	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	}
}

func permissionIf(variable string) bool {
	switch strings.ToLower(variable) {
	case "y", "yes":
		return true
	default:
		return false
	}
}

func installingGit() error {
	if isMacOS() {
		installGitMacOs()
	} else if isWindows() {
		redirectPermission, _ := glogger.ScanQ("Do you want to be redirected to the git download site? [Y/n] ")

		installGitWindows(redirectPermission)
	} else if isDebian() {
		installGitDebian()
	} else if isFedora() {
		installGitFedora()
	} else {
		redirectPermission, _ := glogger.ScanQ("Do you want to be redirected to the git download site? [Y/n] ")

		installGitUnknown(redirectPermission)
	}

	return nil
}

func installGitMacOs() error {
	glogger.Warning("Installing git...")

	cmd := exec.Command("brew", "install", "git")
	return cmd.Run()
}

func installGitWindows(permission string) error {
	if permissionIf(permission) {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://git-scm.com/downloads").Start()
	} else {
		glogger.Info("Downlaod git from: https://git-scm.com/downloads")
	}
	return nil
}

func installGitDebian() error {
	glogger.Warning("Installing git...")
	cmd := exec.Command("sudo", "apt", "install", "git", "-y")
	return cmd.Run()
}

func installGitFedora() error {
	glogger.Warning("Installing git...")

	cmd := exec.Command("sudo", "dnf", "install", "git", "-y")
	return cmd.Run()
}

func installGitUnknown(permission string) error {
	if permissionIf(permission) {
		openBrowser()
	} else {
		glogger.Info("Downlaod git from: https://git-scm.com/downloads")
	}

	return nil
}

func exit() {
	os.Exit(0)
}

// GettingUpdate requests the application update and manages the update process.
func GettingUpdate(permission string) {
	if permissionIf(permission) {
		updateSoftware(false)
	}
}

func updateSoftware(test bool) {
	if checkTheGitPath(test) {
		if checkIfGitIsInstalled() {
			updateWithGit(false)
		} else {
			updateGitNotInstalled()
			if isDebian() || isFedora() || isMacOS() {
				updateWithGit(false)
			}
		}
	} else {
		if !checkIfGitIsInstalled() {
			updateGitNotInstalled()
			createGitPath()
			exit()
		} else {
			createGitPath()
		}
	}

}

func createGitPath() {
	glogger.Warning("Updating wastrap...")
	var cmd *exec.Cmd
	
	cmd = exec.Command("git", "init")
	cmd.Run()
	
	cmd = exec.Command("git", "remote", "add", "origin", "https://github.com/institute-atri/wastrap")
	cmd.Run()
	
	cmd = exec.Command("git", "pull", "origin", "main")
	cmd.Run()
	
	glogger.Done("Update done successfully")
	exit()
}

func updateWithGit(test bool) error {
	glogger.Warning("Updating wastrap...")

	cmd := exec.Command("git", "pull", "origin", "main")
	switch test {
	case true:
		if err := cmd.Run(); err != nil {
			return nil
		}
	case false:
		cmd.Run()
	}

	glogger.Done("Update done successfully")

	return nil
}

func updateGitNotInstalled() {
	glogger.Danger("Git not installed")
	installGitPermission, _ := glogger.ScanQ("Do you want to install git? [Y/n] ")

	if permissionIf(installGitPermission) {
		installingGit()
	}
}
