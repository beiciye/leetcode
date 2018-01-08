package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/count-and-say/description/

func main() {
	fmt.Println(countAndSay(6))
}

func countAndSay(n int) string {

	res := []byte{}
	temp := []byte{}
	var k int
	for i := 0; i < n; i++ {
		temp = []byte{}
		if i == 0 {
			res = []byte("1")
			continue
		}

		k = 1

		for len(res) > 0 {
			if len(res) > 1 && res[0] == res[1] {
				k++
			} else {
				temp = append(temp, []byte(strconv.Itoa(k))...)
				temp = append(temp, res[0])
				k = 1
			}
			res = res[1:]
		}
		res = temp
	}

	return string(res)
}
