package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: appelapsed <input file>")
		return
	}
	custlist, err := os.Open(string(args[1]))
	check(err)
	defer custlist.Close()
	writeTime(start, "Read file")

	outfile, err := os.Create("outfile.csv")
	check(err)
	defer outfile.Close()
	writeTime(start, "Create outfile")

	scanner := bufio.NewScanner(custlist)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), ",")
		outfile.WriteString(names[1] + "," + names[2] + "\n")
	}

	check(scanner.Err())
	writeTime(start, "Wrote data to outfile")
	defer writeTime(start, "Application Finished")
}

func writeTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
