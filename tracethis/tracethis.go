package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("we did not create a trace file! %v\n", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file %v\n", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("we failed to start a trace: %d\n", err)
	}
	trace.Stop()
	AddRandomNumbers()
}

func AddRandomNumbers() {
	f := rand.Intn(100)
	s := rand.Intn(100)
	time.Sleep(10 * time.Second)

	var r = f * s
	fmt.Printf("Result of 2 numbers is %d\n", r)
}
