package main

import (
	"log"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	Fatal
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(file, "Info: ", log.Ldate|log.LUTC|log.Llongfile)
	WarningLogger = log.New(file, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(file, "Fatal: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	writeLog(INFO, "this is a log message!")

	InfoLogger.Println("This is info")
	WarningLogger.Println("This is warning")
	ErrorLogger.Println("This is error")
	FatalLogger.Println("This is fatal")
}

func writeLog(messagetype messageType, message string) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	switch messagetype {
	case INFO:
		logger := log.New(file, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case WARNING:
		logger := log.New(file, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case ERROR:
		logger := log.New(file, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case Fatal:
		logger := log.New(file, "Fatal: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(message)
	}
}
