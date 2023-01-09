package main

import "fmt"

// 错误点
// 1. 忘记更新lastTime
func findPoisonedDuration(timeSeries []int, duration int) int {
	attachDuration := 0
	lastTime := 0
	for idx, i := range timeSeries {
		if idx == 0 {
			lastTime = i
		} else {
			if i-lastTime > duration {
				attachDuration += duration
			} else {
				attachDuration += i - lastTime
			}
			lastTime = i
		}
	}
	attachDuration += duration
	return attachDuration
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := 5
	ret := findPoisonedDuration(a, b)
	fmt.Println(ret)
}
