package main

import "fmt"

func ArrayOfProducts(array []int) []int {
	lp := 0
	rp := len(array) - 1
	product := 1
	var result []int
	for i := 0; i < len(array); i++ {
		for lp != i {
			product = product * array[lp]
			lp++
		}
		for rp != i {
			product = product * array[rp]
			rp--
		}
		fmt.Println(product)

		lp = 0
		rp = len(array) - 1
		result = append(result, product)
		product = 1
	}
	return result
}

func main() {
	var balance = []int{5, 1, 4, 2}

	fmt.Println(ArrayOfProducts(balance))
}
