package main

import (
	"github.com/institute-atri/glogger"
	"github.com/institute-atri/wastrap/cmd"
	"github.com/institute-atri/wastrap/tools/validator"
)

func main() {
	cmd.Execute()
	version := validator.WordpressVersion("https://www.caubr.gov.br/")
	glogger.Done(version)
}
