package main

import (
	"github.com/institute-atri/wastrap/cmd"
	"github.com/institute-atri/wastrap/tools/bruteforce"
)

func main() {
	cmd.Execute()
	bruteforce.Bruteforce("http://localhost/wordpress/", "vendetta")
}
