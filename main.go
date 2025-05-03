package main

import (
	"flag"
	"fmt"
	"runtime"
	"errors"
)


var turbo bool;
func init() {
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
				return "", errors.New("None letter characters deteced")
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

	if len(args) != 1 {
		fmt.Println("Incorrect amount of command line argument value, limited to 1")
		fmt.Println("e.g. `rot26 helloworld`")
		return
	}

	result, err := rot26(args[0])
	if err != nil {
		fmt.Println("Error encounted:", err)
	}
	fmt.Printf("Result: %s", result)
}

