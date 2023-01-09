package main

import (
	"fmt"
	"math"
)

// 错误点:
// 1. 审题，不足三个返回最大值
// 2. 遇到重复的值没考虑
// 3. 错误的将相同的直接continue，没考虑只有三个元素并且有重复值 [1,1,2] 第三大不存在返回2
// 4. 审题: num的取值范围
// 5. math.MinInt32为第三大元素
func thirdMax(nums []int) int {
	first := math.MinInt64
	second := math.MinInt64
	third := math.MinInt64

	for _, i := range nums {
		if i > first {
			third = second
			second = first
			first = i
		} else if i < first && i > second {
			third = second
			second = i
		} else if i < second && i > third {
			third = i
		}
	}
	if len(nums) < 3 && first != math.MinInt64 {
		return first
	}
	if third == math.MinInt64 && first != math.MinInt64 {
		return first
	}
	return third
}

func main() {
	// a := []int{2, 2, 3, 1}
	fmt.Println(math.MinInt32)
	a := []int{1, 1, 2}
	ret := thirdMax(a)
	fmt.Println(ret)
}
