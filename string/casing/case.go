package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "this is a string!"

	stCollection := strings.SplitAfterN(myString, " ", 2)
	for i := range stCollection {
		fmt.Println(stCollection[i])
	}

	fmt.Println(properCase(myString))
}

func properCase(str string) string {
	words := strings.Fields(str)
	smallWord := " a an to the in "

	for i, word := range words {
		if strings.Contains(smallWord, " "+word+" ") {
			words[i] = word
		} else {
			words[i] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
