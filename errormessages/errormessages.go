package main

import (
	"fmt"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;32m%s\033[0m"
)

func main() {
	filename := "test.tx1t"
	showMessage(INFO, fmt.Sprintf("About to open %s", filename))
	showMessage(WARNING, "If the file is not present, this application will fail")

	file, err := os.Open(filename)
	if err != nil {
		showMessage(ERROR, err.Error())
	}
	defer file.Close()

	showMessage(ERROR, "Hello world!")
}

func showMessage(messagetype messageType, message string) {
	switch messagetype {
	case INFO:
		printMessage := fmt.Sprintf("\nInformation: \n%s\n", message)
		fmt.Printf(InfoColor, printMessage)
	case WARNING:
		printMessage := fmt.Sprintf("\nWarning: \n%s\n", message)
		fmt.Printf(WarningColor, printMessage)
	case ERROR:
		printMessage := fmt.Sprintf("\nError: \n%s\n", message)
		fmt.Printf(ErrorColor, printMessage)
	}
}
