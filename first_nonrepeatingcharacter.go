package main

import (
	"fmt"
)

func main() {
	s := "priypdharsshini"
	m := make(map[string]int)
	for _, val := range s {
		m[string(val)]++
	}

	for _, val := range s {
		if m[string(val)] == 1 {
			fmt.Println("The first non repeating character is")
			fmt.Println(string(val))
			break
		}
	}
}
//We go over the string two time which is O(2N) which is O(N) but has space complexity which is O(1)
