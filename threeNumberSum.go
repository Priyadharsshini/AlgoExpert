package main
import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
    sort.Ints(array)
    var result [][]int    
	for i := 0; i < len(array); i++ {
        right := len(array)-1
        left := i+1
        for left < right{
            currentSum := array[i] + array[left] + array[right]
            if currentSum == target{
                result = append(result, []int{array[i], array[left], array[right]})
                left++
                right--
            } else if currentSum < target {
                left ++           
            } else {
                right--
            }         
        }
        
    }
    if len(result) == 0 {
        return [][]int{}
    }
    return result
    }
