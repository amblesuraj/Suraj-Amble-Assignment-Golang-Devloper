package main

import (
	"math"
)

func Rob(nums ...float64) float64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	var value1, value2 float64

	value1 = nums[0]

	value2 = math.Max(value1, nums[1])

	for i := 2; i < n; i++ {
		temp := value2
		value2 = math.Max(value2, value1+nums[i])
		value1 = temp
	}

	return value2

}
