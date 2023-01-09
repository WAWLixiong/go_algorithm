package main

import "fmt"

// 错误点:
// 1. 忽略了结尾是1，需要做额外判断
// 2. 遇到不为1的时候都要将currentCount置为0
func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	currentCount := 0
	for _, i := range nums {
		if i == 1 {
			currentCount++
		} else {
			if currentCount > maxCount {
				maxCount = currentCount
			}
			currentCount = 0
		}
	}
	if nums[len(nums)-1] == 1 {
		if currentCount > maxCount {
			maxCount = currentCount
			currentCount = 0
		}
	}
	return maxCount
}

func main() {
	// a := []int{1, 1, 0, 1, 1, 1}
	a := []int{1, 0}
	ret := findMaxConsecutiveOnes(a)
	fmt.Println(ret)
}
