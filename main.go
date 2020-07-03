package main

import "fmt"

var a = [][]int{
	{0, 1, 2, 3},
	{4, 5, 6, 7},
	{8, 9, 10, 11},
	{1, 5, 7, 8},
}

func main() {
	fmt.Println(sort(a))
}

func sort(slice [][]int) []int {
	var n int = len(slice)
	var sorted []int = slice[0][0 : n-1]

	for i := 0; i < n-1; i++ {
		sorted = append(sorted, slice[i][n-1])
	}

	for m := n - 1; m > 0; m-- {
		sorted = append(sorted, slice[n-1][m])
	}

	for i := n - 1; i > 0; i-- {
		sorted = append(sorted, slice[i][0])
	}

	if len(sorted) < n*n {
		var other [][]int

		for i := 1; i < n-1; i++ {
			other = append(other, slice[i][1:n-1])
		}

		if other != nil {
			sorted = append(sorted, sort(other)...)
		}
	}

	return sorted
}
