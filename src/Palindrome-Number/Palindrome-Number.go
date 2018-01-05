package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(1021))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	i := 1
	for x/i > 0 {
		i *= 10
	}
	result := true
	k := 10
	for i > 10 {

		if x%10 == x*10/i {
			k *= 10

			x = ((x * 10) % i) / 100
			i = i / 100
			continue
		}
		result = false
		break
	}
	return result
}
