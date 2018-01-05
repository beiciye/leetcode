package main

import "fmt"

func main() {
	fmt.Println(strStr("abc", "ab"))

}

//https://leetcode.com/problems/implement-strstr/description/
func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		m := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				m = false
				break
			}
		}
		if m {
			return i
		}
	}

	return -1
}
