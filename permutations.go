package main

import "fmt"

func findPermutations(sr []rune, l int, r int) {
	if l == r {
		fmt.Println(string(sr))
	} else {
		for i := l; i <= r; i++ {
			sr[l], sr[i] = sr[i], sr[l]
			findPermutations(sr, l+1, r)
			sr[l], sr[i] = sr[i], sr[l]
		}
	}
}

func main() {
	s := "ABC"
	sr := []rune(s)
	findPermutations(sr, 0, len(sr)-1)
}
