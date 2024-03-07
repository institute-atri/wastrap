package main

import (
	"fmt"

	"github.com/institute-atri/wastrap/internal/update"
)

func main() {
	fmt.Println("Hello world")
	update.CheckUpdate()
}
