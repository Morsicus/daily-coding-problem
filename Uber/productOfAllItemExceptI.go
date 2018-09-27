/*
Given an array of integers, return a new array
such that each element at index i of the new array
is the product of all the numbers in the original array
except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the
expected output would be [120, 60, 40, 30, 24]. If our input
was [3, 2, 1], the expected output would be [2, 3, 6].
*/

package main

import (
	"fmt"
)

func newArray(arr []int) []int {
	ret := []int{}

	for i := 0; i < len(arr); i++ {
		sum := 1
		for j := 0; j < len(arr); j++ {
			if i != j {
				sum *= arr[j]
			}
			fmt.Println(sum)
		}

		ret = append(ret, sum)
	}

	return ret
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(newArray(arr))
}
