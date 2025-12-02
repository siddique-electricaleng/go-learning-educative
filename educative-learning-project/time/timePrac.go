package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	// sleep for 2 seconds
	time.Sleep(2 * time.Second)
	elapsed := time.Since(t)
	fmt.Println("Elapsed Time: ", elapsed)

	duration := 90 * time.Minute
	fmt.Println("Duration of 90 minutes in nanoseconds: ", duration)

	dhLocation, err := time.LoadLocation("Asia/Dhaka")
	if err != nil {
		fmt.Println("Error loading location: ", err)
		return
	}
	fmt.Println("Location name: ", dhLocation)

	exactTimeNow := time.Now().In(dhLocation)
	fmt.Println("Current time based on location: ", exactTimeNow.Format("01 Jan 2001 15:04"))
	// fmt.Println("Current time based on UTC: ", exactTimeNow)

}
