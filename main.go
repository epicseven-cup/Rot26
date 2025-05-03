package main

import (
	"flag"
	"fmt"
	"runtime"
	"errors"
	"os"
	"bufio"
	"io"
	"strconv"
)


var turbo bool;
func init() {
	flag.BoolVar(&turbo, "t", false, "Enable turbo mode for Rot26 encryption (shorthand)")
	flag.BoolVar(&turbo, "turbo", false, "Enable turbo mode for Rot26 encryption")
}


func rotation(letter rune , direction rune) rune {
	return ((((letter - direction) + 26) % 26) + direction)
}


func rot26(s string) (string, error) {
	if turbo {
		fmt.Println("Turbo mode detected")
		return s, nil 
	} else {
		encrypted := []rune{}
		for _, r := range s {
			var rotatedRune rune;
			if r >= rune('a') && r <= rune('z') {
				// It is lowercase
				rotatedRune = rotation(r, rune('a'))
			} else if r >= rune('A') && r <= rune('Z') {
				// It is uppdercase
				rotatedRune = rotation(r, rune('A'))
			} else {
				return "", errors.New(fmt.Sprintf("None letter character(s) detected, %s", strconv.QuoteRune(r)))
			}
			encrypted = append(encrypted, rotatedRune)
		}
		return string(encrypted), nil
	}
}


func main() {
	fmt.Println("Detecting operating system")
	switch goos := runtime.GOOS; goos {
	case "darwin":
		fmt.Println("OS: darwin")
		fmt.Println("Not supported yet")
		return
	case "linux":
		fmt.Println("OS: linux")
	case "windows":
		fmt.Println("OS: windows")
	default:
		fmt.Println("What is this? Did you really put a monitor and keyboard together and call it a computer?")
		return
	}


	// Parsing the flags
	flag.Parse()
	// Command line args that arn't parsed as flag
	args := flag.Args()
	// Check Pipe
	var input string
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		reader := bufio.NewReader(os.Stdin)
		buffer := make([]byte, 256)
		b := []byte{}

		for {
			n, err := reader.Read(buffer)
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("Error encounted:", err)
			}
			b = append(b, buffer[:n]...)
			buffer = buffer[:n]
		}
		input = string(b)
	} else if len(args) == 1 {
		input = args[0]
	} else {
		fmt.Println("Incorrect amount of command line argument value, limited to 1")
		fmt.Println("e.g. `rot26 helloworld`")
		fmt.Println("e.g. `cat file.txt | rot26`")
		return
	}

	fmt.Printf("Input: %s\n",strconv.Quote(input))

	result, err := rot26(input)
	if err != nil {
		fmt.Println("Error encounted:", err)
	}
	fmt.Printf("Result: %s", result)
}

