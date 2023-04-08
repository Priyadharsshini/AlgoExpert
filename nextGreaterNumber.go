package main

import "fmt"

func findTheNextGreaterNumber(a []int) {
	i := 0
	for i = len(a) - 1; i >= 0; i-- {
		if a[i] > a[i-1] {
			break
		}
		if i == 0 {
			fmt.Println("Not possible")
			break
		}
	}
	start := i
	j := 0

	for j = len(a) - 1; j >= i; j-- {
		if a[i-1] < a[j] {
			a[i-1], a[j] = a[j], a[i-1]
			break
		}
	}

	for i, j := start, end; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println(a)

}

func main() {
	a := []int{5, 3, 4, 9, 7, 6}
	findTheNextGreaterNumber(a)
}
