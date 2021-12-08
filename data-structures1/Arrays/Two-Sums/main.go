package main

import "fmt"

//O(n^2) solution
func twoSum(nums []int, target int) []int {
	var results []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if target-nums[i] == nums[j] {
				results = append(results, i, j)
			}
		}
	}
	return results
}

//using map(hashtable) O(n) time complexity

func twoSums(nums []int, target int) []int {
	var results []int
	mymap := make(map[int]int)

	for i, v := range nums {
		if val, ok := mymap[target-v]; ok {
			results = append(results, val, i)
		}
		mymap[v] = i
	}
	return results
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSums([]int{2, 7, 11, 15}, 9))
}
