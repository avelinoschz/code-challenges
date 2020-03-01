package main

import "fmt"

func compareTriplets(alice, bob []int32) []int32 {

	score := []int32{0, 0}
	for i := 0; i < 3; i++ {
		if alice[i] > bob[i] {
			score[0]++
		} else if bob[i] > alice[i] {
			score[1]++
		}
	}
	return score
}

func main() {
	alice := []int32{5, 6, 7}
	bob := []int32{3, 6, 10}

	score := compareTriplets(alice, bob)
	fmt.Println("output =", score)
}
