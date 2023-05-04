package main

func BestSeat(seats []int) int {
    bestSeat := -1
    maxSpace := 0
    for left, right := 0,1; left < len(seats)-1 && right < len(seats)-1;{
        for seats[right] != 1 {
            right++
        }
        if right-left-1 > maxSpace {
            maxSpace = right-left-1
            bestSeat = (right+left)/2
        }
        left = right
        right++
    }
	return bestSeat
}
