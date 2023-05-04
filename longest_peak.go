package main

func LongestPeak(array []int) int {
    if len(array) < 3{
        return 0
    }
    Longest := 0
    for i := 1; i < len(array)-1; {
        if array[i] > array[i-1] && array[i] > array[i+1]{
            left := i-2
            for left >=0 && array[left] < array[left+1]{
                left--
            }
            right := i+2
            for right <= len(array)-1 && array[right] < array[right-1]{
                right++
            }
            currentLongest := right - left - 1
            if currentLongest > Longest{
                    Longest = currentLongest
                   
            }
            i = right           
            }else {
                i++
        }
    }
	
	return Longest
}
