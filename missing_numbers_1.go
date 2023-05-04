package main

func MissingNumbers(nums []int) []int {
    length := len(nums) + 2
    sum := 0
    curSum := 0
    left := 0
    right := 0
    tobeleft := 0
    toberight := 0
    var result []int
    for i := 1; i <= length; i++ {
        sum += i
    }
    for i := range nums {
        curSum += nums[i] 
    }
    avg := (sum - curSum)/2
    for i := 0; i < len(nums); i++ {
        if nums[i] <= avg {
            left += nums[i]
        } else {
            right += nums[i]
        }
    }
    for i := 1; i <= avg; i++ {
        tobeleft += i
    }
    for i := avg + 1; i <= length; i++ {
        toberight += i
    }
    missingleft := tobeleft - left
    result = append(result, missingleft)
    missingright := toberight - right
    result = append(result, missingright)

    return result
}
