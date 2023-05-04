package main

func ZeroSumSubarray(nums []int) bool {
    curSum := 0
    allSum := make(map[int]int)
    for i := 0 ; i< len(nums); i++ {
        curSum += nums[i]
        if curSum == 0{
            return true
        } 
        if _, ok := allSum[curSum];ok {
            return true
        } else {
            allSum[curSum] = nums[i]
        }
    }
	return false
}
