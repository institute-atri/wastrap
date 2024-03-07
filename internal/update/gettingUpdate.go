package update

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func checkTheGitPath() bool {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("The program is damaged,: https://github.com/institute-atri/wastrap")
	}

	gitDir := filepath.Join(currentDir, ".git")
	_, err = os.Stat(gitDir)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		fmt.Println("Error when checking if the project is using git")
		return false
	}
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
	var err error
	url := "https://git-scm.com/downloads"

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		fmt.Println(err)
	}
}

func permissionIf(variable string) bool {
	if variable == "y" || variable == "Y" || variable == "yes" || variable == "YES" || variable == "Yes" {
		return true
	} else {
		return false
	}
}

func installingGit() {
	url := "https://git-scm.com/downloads"

	if isMacOS() {
		fmt.Println("Installing git...")

		cmd := exec.Command("brew", "install", "git")
		cmd.Run()

	} else if isWindows() {
		var redirectPermission string

		fmt.Print("Do you want to be redirected to the git download site? [Y/n] ")
		fmt.Scan(&redirectPermission)

		if permissionIf(redirectPermission) {
			exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		} else {
			fmt.Println("Downlaod git from: https://git-scm.com/downloads")
		}

	} else if isDebian() {
		fmt.Println("Installing git...")

		cmd := exec.Command("sudo", "apt", "install", "git", "-y")
		cmd.Run()

	} else if isFedora() {
		fmt.Printf("Installing git...")

		cmd := exec.Command("sudo", "dnf", "install", "git", "-y")
		cmd.Run()

	} else {
		var redirectPermission string

		fmt.Print("Do you want to be redirected to the git download site? [Y/n] ")
		fmt.Scan(&redirectPermission)

		if permissionIf(redirectPermission) {
			openBrowser()
		} else {
			fmt.Println("Downlaod git from: https://git-scm.com/downloads")
		}
	}
}

func GettingUpdate() {
	var updateWastrapPermission string

	fmt.Print("Do you want to update wastrap [Y/n] ")
	fmt.Scan(&updateWastrapPermission)

	if permissionIf(updateWastrapPermission) {
		if checkTheGitPath() {
			if checkIfGitIsInstalled() {
				fmt.Println("Updating wastrap")

				cmd := exec.Command("git", "pull", "origin", "main")
				cmd.Run()

				println("Update done successfully")
			} else {
				var installGitPermission string

				fmt.Print("Do you want to install git? [Y/n] ")
				fmt.Scan(&installGitPermission)

				if permissionIf(installGitPermission) {
					installingGit()
				}

				if isDebian() || isFedora() || isMacOS() {
					fmt.Println("Updating wastrap")

					cmd := exec.Command("git", "pull", "origin", "main")
					cmd.Run()

					println("Update done successfully")
				}
			}
		} else {
			if checkIfGitIsInstalled() {
				fmt.Println("Updating wastrap")
				var cmd *exec.Cmd
				cmd = exec.Command("git", "init")
				cmd.Run()
				cmd = exec.Command("git", "remote", "add", "origin", "https://github.com/institute-atri/wastrap")
				cmd.Run()
				cmd = exec.Command("git", "pull", "origin", "main")
				cmd.Run()
				println("Update done successfully")
			} else {
				var installGitPermission string

				fmt.Print("Do you want to install git? [Y/n] ")
				fmt.Scan(&installGitPermission)

				if permissionIf(installGitPermission) {
					installingGit()
				}

				if isDebian() || isFedora() || isMacOS() {
					fmt.Println("Updating wastrap")
					var cmd *exec.Cmd
					cmd = exec.Command("git", "init")
					cmd.Run()
					cmd = exec.Command("git", "remote", "add", "origin", "https://github.com/institute-atri/wastrap")
					cmd.Run()
					cmd = exec.Command("git", "pull", "origin", "main")
					cmd.Run()
					println("Update done successfully")
				}
			}
		}
	}
}
