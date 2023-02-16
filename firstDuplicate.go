package main

func FirstDuplicateValue(array []int) int {
    m := make(map[int]int)
    for _, value :=  range array{
        _, ok := m[value]
        if ok {
            return value
        } else {
            m[value] = 0
        }
    }
	return -1
}
