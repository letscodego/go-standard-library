package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan(){
			if scanner.Text() =="/quit"{
				fmt.Println("Quitting")
				os.Exit(3)
			}else{
				fmt.Println("You typed " + scanner.Text())
			}
		}
	*/
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
