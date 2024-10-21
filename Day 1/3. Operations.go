package main

import "fmt"

func main() {
	x,y := 50,3

	fmt.Print("Addition: ", x+y, "\n",
				"Subtraction: ", x-y, "\n",
				"multiplication: ",x*y, "\n",
				"division: ", x/y, "\n",
				"modulus: ", x%y, "\n" )
}