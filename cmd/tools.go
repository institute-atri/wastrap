package cmd

import (
	"fmt"
	"time"
)

func formatMessage(text string, style int) string {
	switch style {
	case 1:
		return fmt.Sprintf("\033[1;31m%s\033[0m", text)
	case 2:
		return fmt.Sprintf("\033[1;34m%s\033[0m", text)
	default:
		return text
	}
}

func loading() {
	fmt.Print("Loading... ")
	for i := 0; i < 5; i++ {

		fmt.Print("|")
		time.Sleep(150 * time.Millisecond)
		fmt.Print("\b")
		fmt.Print("/")
		time.Sleep(150 * time.Millisecond)
		fmt.Print("\b")
		fmt.Print("-")
		time.Sleep(150 * time.Millisecond)
		fmt.Print("\b")
		fmt.Print("\\")
		time.Sleep(150 * time.Millisecond)
		fmt.Print("\b")
	}
	fmt.Println()
}
