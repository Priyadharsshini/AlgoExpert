package main

import "sort"
func MissingNumbers(nums []int) []int {
    hm := make(map[int]int)
    var result []int
    length := len(nums)+2
    
    for i := 1; i <= length; i++ {
        hm[i] = 1
    }
    for _, val := range nums{
        if _, ok := hm[val]; ok {
            hm[val]++
        }
    }
    for val, count := range hm{
        if count == 1{
            result = append(result, val)
        }
    }
    sort.Ints(result)
    return result
}


