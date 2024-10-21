package main

import "fmt"

func main() {

	for i := 1 ; i <= 5 ; i++ { // for loop
		fmt.Println(i)
	}

	i := 1

	for i<=10 { // for as while
		fmt.Print(i)
		i++
	}

	for i:= 1; i < 5 ; i++ { // nested loops

		for j := 1; j<i; j++ {
			fmt.Printf("*")
		}

		fmt.Println()
	}
}