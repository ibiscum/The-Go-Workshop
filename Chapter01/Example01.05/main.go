package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	t := time.Now()

	// Defining duration
	d := (60 * time.Second)

	// Calling Truncate() method
	trunc := t.Truncate(d)
	return false, "info", trunc
}

func main() {
	// Type only
	var start, middle, end float32
	fmt.Println(start, middle, end)
	// Initial value mixed type
	var name, left, right, top, bottom = "one", 1, 1.5, 2, 2.5
	fmt.Println(name, left, right, top, bottom)
	// works with functions also
	var Debug, LogLevel, startUpTime = getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}
