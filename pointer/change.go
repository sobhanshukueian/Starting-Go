package main

import "fmt"

func modifyValue(val *int) {
	*val = 10 // Update the value at the memory address pointed to by val
}

func main() {
	x := 5
	fmt.Println("Before:", x) // Before: 5

	modifyValue(&x)          // Pass the address of x to the function
	fmt.Println("After:", x) // After: 10
}
