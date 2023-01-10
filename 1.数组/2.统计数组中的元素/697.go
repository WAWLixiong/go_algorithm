package main

import (
	"fmt"
	"math"
)

func findShortestSubArray(nums []int) int {
	book := make(map[int]int, len(nums))
	for _, i := range nums {
		book[i]++
	}

	maxTimes := 0
	sameKey := make(map[int]int, 0)
	for _, v := range book {
		if v > maxTimes {
			maxTimes = v
		}
	}
	for k, v := range book {
		if v == maxTimes {
			sameKey[k] = 0
		}
	}

	pathBook := make(map[int]int)
	for _, i := range nums {
		for k := range sameKey {
			vv := pathBook[i]
			if i == k {
				pathBook[i]++
			}
			if i == k || (vv > 0 && vv != maxTimes) {
				sameKey[k]++
			}
		}
	}

	minPath := math.MaxInt
	for _, v := range sameKey {
		if v < minPath {
			minPath = v
		}
	}
	return minPath
}

func main() {
	a := []int{1, 2, 2, 3, 1, 4, 2}
	ret := findShortestSubArray(a)
	fmt.Println(ret)
}
