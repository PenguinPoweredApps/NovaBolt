package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/atotto/clipboard"
)

const (
	lowerCase = "abcdefghijklmnopqrstuvwxyz"
	upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	special   = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func main() {
	// Define CLI flags
	length := flag.Int("l", 16, "Length of the password")
	useUpper := flag.Bool("u", true, "Include uppercase letters")
	useNumbers := flag.Bool("n", true, "Include numbers")
	useSpecial := flag.Bool("s", true, "Include special characters")
	timeout := flag.Int("t", 30, "Seconds before clipboard is cleared")

	flag.Parse()

	// Generate the password
	password, err := generatePassword(*length, *useUpper, *useNumbers, *useSpecial)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Copy to clipboard
	err = clipboard.WriteAll(password)
	if err != nil {
		fmt.Printf("Error copying to clipboard: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("🔑 Password of length %d generated and copied to clipboard!\n", *length)
	fmt.Printf("⏳ It will be cleared in %d seconds...\n", *timeout)

	// Block and clear
	clearClipboard(*timeout)
}

func generatePassword(length int, upper, num, spec bool) (string, error) {
	charPool := lowerCase
	if upper {
		charPool += upperCase
	}
	if num {
		charPool += numbers
	}
	if spec {
		charPool += special
	}

	if len(charPool) == 0 {
		return "", fmt.Errorf("no character sets selected")
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charPool))))
		if err != nil {
			return "", err
		}
		password[i] = charPool[index.Int64()]
	}

	return string(password), nil
}

func clearClipboard(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
	clipboard.WriteAll("")
	fmt.Println("🧹 Clipboard cleared.")
}
