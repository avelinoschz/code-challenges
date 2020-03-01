package main

import "fmt"

func firstFactorial(num int) int {

	// code goes here
	// Note: feel free to modify the return type of this function

	fmt.Println("input =", num)

	aux := 1

	for i := 1; i <= num; i++ {
		aux *= i
	}

	return aux
}

func firstFactorialTOP(num int) int {
	if num <= 1 {
		return 1
	}
	return num * firstFactorialTOP(num-1)

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you

	result := firstFactorialTOP(4)
	fmt.Println("output =", result)

	// Test case #1
	// input 4
	// output 24
}
