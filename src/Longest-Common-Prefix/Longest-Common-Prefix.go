package main

import "fmt"

//https://leetcode.com/problems/longest-common-prefix/

func main() {
	fmt.Println(longestCommonPrefix([]string{"c", "c"}))
}

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	i := 0
	m := true
	for m {
		for _, val := range strs {
			if len(val) == 0 {
				m = false
				break
			}

			if i == len(val) {
				m = false
				break
			}

			if val[i] == 0 {
				m = false
				break
			}

			if val[i] == strs[0][i] {
				continue
			} else {
				m = false
				break
			}
		}
		i++
	}
	return strs[0][:i-1]
}
