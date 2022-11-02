package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "/help" {
		fmt.Println("Usage: filemaker <input file>")
	} else {
		fmt.Println("How would you like to see the text?")
		fmt.Println("1: ALL CAPS")
		fmt.Println("2: Title Case")
		fmt.Println("3: lower case")

		var option int
		_, err := fmt.Scanf("%d", &option)
		if err != nil {
			fmt.Println(err)
		}

		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			switch option {
			case 1:
				fmt.Println(strings.ToUpper(scanner.Text()))
			case 2:
				fmt.Println(strings.Title(scanner.Text()))
				fmt.Println(cases.Title(language.English).String(scanner.Text()))
			case 3:
				fmt.Println(strings.ToLower(scanner.Text()))
			}
		}
	}
}
