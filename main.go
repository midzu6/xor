package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"xor/sipher"
)

var mode = flag.String("mode", "cipher", "Set 'sipher' or 'decipher'. Default is cipher")
var secretKey = flag.String("secret", "", "your secret key must contain at leat 1 symbol")

func userInput(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v, enter command again\n", err)
			continue
		}
		return strings.TrimSpace(input)
	}
}

func printError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
}

func main() {

	flag.Parse()

	switch *mode {
	case "cipher":
		plaintext := userInput("Enter text for cipher:\n> ")
		encodedOutput, err := sipherer.Cipher(plaintext, *secretKey)
		if err != nil {
			printError(err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", encodedOutput)
	case "decipher":
		plaintext := userInput("Enter text for decipher:\n> ")
		decodedOutput, err := sipherer.Decipher(plaintext, *secretKey)
		if err != nil {
			printError(err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", decodedOutput)
	default:
		fmt.Println("invalid mode. Use 'cipher' or 'decipher'.")
		os.Exit(1)
	}
}
