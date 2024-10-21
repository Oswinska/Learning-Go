package main

import "fmt"

func main() {

	x := 10

	fmt.Println(x, &x)

	changeValue(x) // Won't Change x
	fmt.Println(x)

	changeValuePointer(&x) // Points to the memory of x defined in main, thus changes X
	fmt.Println(x)
}

func changeValue(x int) {

	x = 7;
}

func changeValuePointer(x *int) {

	*x = 7;
}