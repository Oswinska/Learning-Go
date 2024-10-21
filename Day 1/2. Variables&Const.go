package main

import "fmt"

func values() {
	var a int = 600
	var b float32 = 6.9
	const bombotron = 256

	x,y,z := 1,2,3 // short variable declaration

	var isTrue bool = true // Don't forget var everywhere besides ^
	isBool := false // shorter way to declare a boolean

	fmt.Println(a,b,bombotron,x,y,z, "Bool: ", isTrue, isBool) // painless chaining
	fmt.Print(
		isTrue && isBool,
		isTrue || isBool,
		!isTrue, "\n" )
} 

func main() {
	values()
}
