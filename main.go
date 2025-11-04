package main

import "fmt"

func Sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	s := 0
	for _, i := range arr {
		s += i
	}
	return s
}

func main() {
	a := [][]int{{1, 2, 3}, {}, {0}}
	for _, arr := range a {
		fmt.Println(Sum(arr))
	}
}
