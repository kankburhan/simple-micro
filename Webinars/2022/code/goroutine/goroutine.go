package main

// OMIT
import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// START OMIT
	// Loop Go Routine
	runtime.GOMAXPROCS(4)
	start := time.Now()
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	totalEksekusi := time.Since(start)
	fmt.Println("Total Eksekusi: " + totalEksekusi.String())
	time.Sleep(time.Second)
	// END OMIT
}
