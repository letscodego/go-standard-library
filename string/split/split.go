package main

import (
	"fmt"
	"strings"
)

func main() {
	mystring := "this is a string!"

	stCollection := strings.SplitAfterN(mystring, " ", 2)
	for i := range stCollection {
		fmt.Println(stCollection[i])
	}
}
