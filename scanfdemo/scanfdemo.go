package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("What's your name?")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello %s! Nice to meet you!\n", name)

	fmt.Printf("|%7.2f|%7.2f|%7.2f|%7.2f|\n", 233.5445, 52.5877, 66.8797, 4556.4)
	fmt.Printf("|%7.2f|%7.2f|%7.2f|%7.2f|\n", 23.5445, 52.58, 6.897, 46.4)

	fmt.Printf("|%-7.2f|%-7.2f|%-7.2f|%-7.2f|\n", 233.5445, 52.5877, 66.8797, 4556.4)
	fmt.Printf("|%-7.2f|%-7.2f|%-7.2f|%-7.2f|\n", 23.5445, 52.58, 6.897, 46.4)

	output := fmt.Sprintf("|%-7.2f|%-7.2f|%-7.2f|%-7.2f|\n", 23.5445, 52.58, 6.897, 46.4)
	print(output)
}
