package bruteforce

import (
	"os"
	"testing"

	"github.com/institute-atri/glogger"
)

func TestBruteforce(t *testing.T) {
	file, err := os.Create("wordtest.txt")
	if err != nil {
		glogger.Fatal("Error when creating the file", err)
	}
	defer file.Close()

	_, err = file.WriteString("password123")
	if err != nil {
		glogger.Fatal("Error when write in the file", err)
	}


	Bruteforce("http://exampleATRI.com", "admin", "./wordtest.txt")
	Bruteforce("http://exampleATRI.com", "admin", "./wordtest.txt")
	Bruteforce("http://exampleATRI.com", "admin", "./fail.tsx")


	err = os.Remove("wordtest.txt")
	if err != nil {
		glogger.Fatal("Error when removing the file", err)
	}
}