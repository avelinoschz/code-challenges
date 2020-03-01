package main

import "fmt"

func firstReverse(str string) string {

	reversed := ""
	for i := len(str); i > 0; i-- {
		reversed += string(str[i-1])
	}

	return reversed
}

func main() {
	result := firstReverse("coderbyte")
	fmt.Println("output =", result)
}
