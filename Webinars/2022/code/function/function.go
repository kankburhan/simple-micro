// *build ignore,OMIT
package main

import "fmt"

// STARTMAIN1 OMIT

// function as value
func FuncAsValue(a, b int) int {
	return a * b
}

func FuncAsClosure() func() int {
	i := 0
	return func() int {
		i = i + 1
		return i
	}
}

func main() {
	valofClosure := FuncAsClosure()
	fmt.Println(FuncAsValue(4, 4))
	fmt.Println(valofClosure(), valofClosure(), valofClosure())
}

// STOPMAIN1 OMIT
