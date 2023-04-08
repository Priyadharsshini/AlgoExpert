
package main

import (
	"fmt"
)

func main() {
	s := "priypdharsshini"
	foundDuplicate := false
	for i := range s {
		foundDuplicate = false
		for j := range s {
			if s[i] == s[j] && i != j {
				foundDuplicate = true
				break
			}
		}
		if foundDuplicate == false {
			fmt.Println("The first non repeating character is")
			fmt.Println(string(s[i]))
			return
		}

	}
}

//O(N^2) since two nested for loops going over the entire string
