package main

import (
	"fmt"
	"strings"
)

func longestSubString(s string) string {
	var temp string
	var result string
	for i := 0; i < len(s); i++ {
		temp = ""
		for j := i; j < len(s); j++ {
			if !strings.Contains(temp, string(s[j])) {
				temp += string(s[j])
			} else {
				break
			}

			if len(temp) > len(result) {
				result = temp
			}
		}
	}

	return result
}

func main() {
	s := "abcabcdd"
	result := longestSubString(s)
	fmt.Println(result)
}
