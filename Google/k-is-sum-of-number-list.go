package main

import (
	"fmt"
)

/*
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
*/

func isSumOfNumber(s []int, k int) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i]+s[j] == k {
				return true
			}
		}
	}

	return false
}

func main() {
	s := []int{10, 15, 3, 7}
	s1 := []int{11, 2, 5, 8, 10}
	s2 := []int{1, 2, 3, 7, 10}
	k := 17

	fmt.Println(isSumOfNumber(s, k))
	fmt.Println(isSumOfNumber(s1, k))
	fmt.Println(isSumOfNumber(s2, k))
}
