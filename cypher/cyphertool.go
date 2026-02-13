package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() (toEncrypt bool, encoding string, message string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Select operation (1/2): ")
		fmt.Println("1. Encrypt.")
		fmt.Println("2. Decrypt.")
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err == nil && (num == 1 || num == 2) {
			toEncrypt = (num == 1)
			break
		}
		fmt.Println("Invalid choice. Try again! ")
	}
	for {
		fmt.Println("\nSelect cypher (1/2): ")
		fmt.Println("1. ROT13.")
		fmt.Println("2. Reverse.")
		fmt.Println("3. ShiftCap.")
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err == nil {
			if num == 1 {
				encoding = "rot13"
				break
			} else if num == 2 {
				encoding = "reverse"
				break
			} else if num == 3 {
				encoding = "shiftcap"
				break
			}
			fmt.Println("Invalid choice. Try again!")
		}
	}
	fmt.Println("\nEnter the message: ")
	scanner.Scan()
	message = scanner.Text()
	message = strings.TrimSpace(message)
	return toEncrypt, encoding, message
}

func encrypt_rot13(s string) string {
	var builder strings.Builder
	builder.Grow(len(s))
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char = 'a' + (char-'a'+13)%26
		} else if char >= 'A' && char <= 'Z' {
			char = 'A' + (char-'A'+13)%26
		}
		builder.WriteRune(char)
	}
	return builder.String()
}

func decrypt_rot13(s string) string {
	return encrypt_rot13(s)
}

func encrypt_reverse(s string) string {
	var builder strings.Builder
	builder.Grow(len(s))
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char = 'a' + ('z' - char)
		} else if char >= 'A' && char <= 'Z' {
			char = 'A' + ('Z' - char)
		}
		builder.WriteRune(char)
	}
	return builder.String()
}

func decrypt_reverse(s string) string {
	return encrypt_reverse(s)
}

func encrypt_shiftcap(s string) string {
	var builder strings.Builder
	builder.Grow(len(s))
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char = ('a' + (char-'a'+5)%26) - 32
		} else if char >= 'A' && char <= 'Z' {
			char = ('A' + (char-'A'+5)%26) + 32
		}
		builder.WriteRune(char)
	}
	return builder.String()
}

func decrypt_shiftcap(s string) string {
	var builder strings.Builder
	builder.Grow(len(s))
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char = 'a' + (char-'a'-5+26)%26 - 32
		} else if char >= 'A' && char <= 'Z' {
			char = 'A' + (char-'A'-5+26)%26 + 32
		}
		builder.WriteRune(char)
	}
	return builder.String()
}

func main() {
	fmt.Println("\nWelcome to the Cypher Tool!\n")
	toEncrypt, encoding, msg := getInput()
	if toEncrypt == true {
		if encoding == "rot13" {
			fmt.Println("\nEncrypted message using rot13: ")
			fmt.Println(encrypt_rot13(msg))
		} else if encoding == "reverse" {
			fmt.Println("\nEncrypted message using reverse: ")
			fmt.Println(encrypt_reverse(msg))
		} else {
			fmt.Println("\nEncrypted message using shiftcap: ")
			fmt.Println(encrypt_shiftcap(msg))
		}
	} else {
		if encoding == "rot13" {
			fmt.Println("\nDecrypted message using rot13: ")
			fmt.Println(decrypt_rot13(msg))
		} else if encoding == "reverse" {
			fmt.Println("\nDecrypted message using reverse: ")
			fmt.Println(decrypt_reverse(msg))
		} else {
			fmt.Println("\nDecrypted message using shiftcap: ")
			fmt.Println(decrypt_shiftcap(msg))
		}
	}
}
