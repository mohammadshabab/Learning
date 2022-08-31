package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	_ = start
	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
		return
	}

	dateString := os.Args[1]

	// date only so time will be 00:00
	d, err := time.Parse("02 January 2006", dateString)
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Time: ", d.Day(), d.Month(), d.Year())
	}

	d, err = time.Parse("02 January 2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Date: ", d.Day(), d.Month(), d.Year())
		fmt.Println("Time: ", d.Hour(), d.Minute())
	}

	d, err = time.Parse("02-01-2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Date: ", d.Day(), d.Month(), d.Year())
		fmt.Println("Time: ", d.Hour(), d.Minute())
	}

	d, err = time.Parse("15:04", dateString)
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Date: ", d.Day(), d.Month(), d.Year())
		fmt.Println("Time: ", d.Hour(), d.Minute())
	}

	t := time.Now().Unix()
	fmt.Println("Epoch time ", t)

	d = time.Unix(t, 0)
	fmt.Println("Date: ", d.Day(), d.Month(), d.Year())
	fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())

	duration := time.Since(start)
	fmt.Println("Execution time: ", duration)

	d, err = time.Parse("02 January 2006 15:04 MST", dateString)
	if err != nil {
		fmt.Println(err)
	}
	nd := time.Now()
	loc, _ := time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", d.In(loc))

	nloc, _ := time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", nd.In(nloc))
}
