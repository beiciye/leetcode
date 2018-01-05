package main

import (
	"fmt"
)

// https://leetcode.com/problems/reverse-integer/description/

func main() {
	w := 21
	fmt.Println(reverse(w))
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	positive := true
	if x < 0 {
		positive = false
		x = -x
	}

	i := 1
	var result int
	for x/i > 0 {
		i = i * 10
		result = result*10 + (x%i)*10/i
	}

	if result >= (1 << 31) {
		return 0
	}
	if !positive {
		result = -result
	}
	return result

}
