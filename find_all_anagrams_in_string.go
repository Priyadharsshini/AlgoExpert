package main

import (
	"fmt"
	"reflect"
)

func findAnagrams(s string, p string) []int {
	patterMap := make(map[string]int)
	originalMap := make(map[string]int)
	finalResult := make([]int, 0)
	k := 0

	for _, val := range p {
		patterMap[string(val)]++
	}
	i := 0

	for ; i < len(s); i++ {
		originalMap[string(s[i])]++
		k++
		if k == len(p) {
			res1 := reflect.DeepEqual(patterMap, originalMap)
			if res1 {
				finalResult = append(finalResult, i-len(p)+1)
			}
			i = i - len(p) + 1
			k = 0
			for key := range originalMap {
				delete(originalMap, key)
			}

		}

	}

	return finalResult
}
func main() {
	fmt.Println(findAnagrams("abab", "ab"))

}
