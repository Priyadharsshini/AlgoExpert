package main

func IsMonotonic(array []int) bool {
    if len(array) == 0 || len(array) ==1 {
        return true
    }
    k := 0
    l := 1
    for array[k] == array[l] && k < len(array){
        if l == len(array)-1{
            return true
        }
        k++
        l++
    }
    if array[k] > array[l]{
        for i := 1; i < len(array)-1; i++{
            if array[i] < array[i+1]{
                return false
            }
        }
    }else if array[k] < array[l]{
        for i := 1; i < len(array)-1; i++ {
            if array[i] > array[i+1]{
                return false
            }
        }
    } 
	return true
}
