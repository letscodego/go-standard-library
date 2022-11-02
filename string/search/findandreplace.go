package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	myString := "I really enjoy the GO language, It's one of my favorites!"

	if len(os.Args) < 2 {
		fmt.Println("Enter search term")
	} else {
		searchTerm := os.Args[1]
		//result := strings.HasPrefix(myString, searchTerm)
		result := strings.Contains(myString, searchTerm)
		//results := strings.Replace(myString, "I", "i", -1)

		if result {
			fmt.Printf("the sample string contains %s\n", searchTerm)
		} else {
			fmt.Printf("the sample string dose not contain %s\n", searchTerm)
		}
	}
}
