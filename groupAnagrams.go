package main

import (
       "sort")

func GroupAnagrams(words []string) [][]string {
    hm := make(map[string][]string)
	  for _, word := range words{
        char := []rune(word)
      //The function compares the characters at the two indices i and j using the less than operator <. 
      //If the character at index i is less than the character at index j, the function returns true, indicating that the element at index i 
      //should come before the element at index j in the sorted slice.The sort.Slice function then uses the function provided as the second argument 
      //to determine the order of the elements in the slice and sorts the slice accordingly.
        sort.Slice(char, func(i,j int)bool{
            return char[i]<char[j]
        })
        anagram := string(char)
        if _, ok := hm[anagram]; ok {
            hm[anagram] = append(hm[anagram], word)
        } else {
             hm[anagram] = append(hm[anagram], word)
        }
        }
  //Initial length is 0 and the capacity is the initial max capacity, atleast that many elements it should hold.
    result := make([][]string, 0, len(hm))
	for _, value := range hm {
		result = append(result, value)
	}

	return result
   
    }
	
