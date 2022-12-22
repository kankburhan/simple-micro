// *build ignore,OMIT
package main

import "fmt"

// STARTMAIN1 OMIT

func main() {
	// short variabel declare
	var kd, al string = "Klikdokter", "Alchemy"
	product, price := "Air Mineral", 3000

	fmt.Println(kd, " x ", al)
	fmt.Println(product, "-", price)
}

// STOPMAIN1 OMIT
