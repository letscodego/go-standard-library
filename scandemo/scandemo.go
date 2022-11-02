package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("What's your name")
	//inputs, _ := fmt.Scanf("%s", &name)
	inputs, _ := fmt.Scanf("%q", &name) //to read multiple words with quests

	switch inputs {
	case 0:
		fmt.Printf("You must enter a name!\n")
	case 1:
		fmt.Printf("Hello %s! Nice to meet you\n", name)
	}
}
