package main

import (
	"fmt"
	"time"
)

/**
Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
https://golang.org/pkg/time/
*/

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
	fmt.Println(now.Format(time.RFC822))
	fmt.Println(now.Format(time.Kitchen))
	fmt.Println(now.Format(time.UnixDate))
	fmt.Println(now.Format(time.RFC3339))

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02Z+07"
	// layout := "2023-01-02"
	layout := time.RFC3339
	parse, _ := time.Parse(layout, "2023-03-11Z+07")
	fmt.Print(parse)
}
