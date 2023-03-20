package main

import "math"
import "sort"

func SmallestDifference(array1, array2 []int) []int {
    sort.Ints(array1)
    sort.Ints(array2)
    var result = make([]int, 2)
    diff := int(math.Abs(math.Abs(float64(array1[0])- float64(array2[0]))))
    for i, j := 0, 0 ; i < len(array1) && j < len(array2);{
        currdiff := int(math.Abs(float64(array1[i])- float64(array2[j])))
        if currdiff == 0 {
            result[0] = array1[i]
            result[1] = array2[j]
            return result
        }
        if currdiff < diff {
            diff = currdiff
            result[0] = array1[i]
            result[1] = array2[j]
        } else {
            if array1[i] < array2[j]{
                i++
            } else {
                j++
            }
        }        
    }
    return result    
}
