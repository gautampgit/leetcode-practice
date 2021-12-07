package main

import (
	"fmt"
	"math"
)

//O(n^2) solution
func maxSubArray1(nums []int) {
	max_sum := math.MinInt

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if max_sum < sum {
				max_sum = sum
			}
		}
	}
	fmt.Println(max_sum)
}

//Kadane algorithm O(n) solution
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

//effecient kadane algorithm
func maxSubArray2(nums []int) int {
	max := nums[0]
	sum := 0
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
func main() {
	fmt.Println(maxSubArray([]int{-2, 2, 5, -11, 6}))
}
