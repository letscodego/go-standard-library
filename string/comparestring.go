package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "this is a string!"
	string2 := "This is a string!"
	if strings.Compare(string1, string2) == 0 {
		fmt.Println("the strings are identical")
	} else {
		fmt.Println("the strings are not identical")
	}

	st := []string{"Larry", "Mot", "Dell"}
	for _, stooge := range st {
		fmt.Println(strings.Compare("Larry", stooge))
	}
}
