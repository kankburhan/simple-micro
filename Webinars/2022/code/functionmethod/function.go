// *build ignore,OMIT
package main

import "fmt"

// STARTMAIN1 OMIT

type Persegi struct {
	panjang, lebar int
}

// func (receiver) funcName() return
func (persegi Persegi) Luas() int {
	return persegi.panjang * persegi.lebar
}

func main() {
	persegi := Persegi{panjang: 4, lebar: 4}
	luasPersegi := persegi.Luas()
	fmt.Println("Luas Persegi : ", luasPersegi)
}

// STOPMAIN1 OMIT
