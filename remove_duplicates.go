package main

import "fmt"

func main() {
	duplicatedString := "apple"
	storageMap := make(map[string]int)
	var unDuplicatedString string
	for _, key := range duplicatedString {
		if _, ok := storageMap[string(key)]; ok {
			storageMap[string(key)]++
		} else {
			storageMap[string(key)] = 1
		}

	}

	for key, value := range storageMap {
		if value == 1 {
			unDuplicatedString += key
		}
	}

	fmt.Println(unDuplicatedString)
}
