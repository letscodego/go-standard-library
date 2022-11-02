package main

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
	layoutEU  = "2 January, 2006"
)

func main() {
	t := time.Now()
	fmt.Print(t, "\n")
	fmt.Print(t.Year(), "\n")

	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format(time.Kitchen))

	//Mon Jan 2 15:04:05 MST 2006  based on go document
	fmt.Println(t.Format("Mon Jan 2 15:04:05 MST 2006"))
	fmt.Println(t.Format("Monday, January 2 in the year 2006"))

	startDate := time.Date(2017, 04, 04, 01, 04, 04, 04, time.UTC)
	fmt.Println(startDate.Format("Monday, Jan 2 at 15:04:04 2006"))

	fmt.Println(startDate.Format(layoutISO))
	fmt.Println(startDate.Format(layoutUS))
	fmt.Println(startDate.Format(layoutEU))

	today := time.Now()
	future := today.AddDate(0, 6, 0)
	fmt.Printf("In six months it will be %v\n", future.Format("Monday, January 2"))
}
