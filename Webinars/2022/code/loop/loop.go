package main

import (
	"fmt"
	"time"
)

func main() {
	// Loop without Go Routine
	start := time.Now()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	totalEksekusi := time.Since(start)

	fmt.Println("Total Eksekusi: " + totalEksekusi.String())
}
