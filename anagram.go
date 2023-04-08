package main

import "fmt"

func isAnagram(s string, t string) {
	m := make(map[string]int)
	if len(s) != len(t) {
		fmt.Println("Not an anagram")
		return
	}
	for val := range s {
		m[string(val)]++
	}
	for val := range t {
		m[string(val)]--
	}

	for _, val := range m {
		if val != 0 {
			fmt.Println("Not an anagram")
			return
		}
	}

	fmt.Println("It is an anagram")

}

func main() {
	s := "apple"
	t := "elapp"
	isAnagram(s, t)
}
