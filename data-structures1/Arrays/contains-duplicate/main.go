/* return true if the array contains a duplicate else return false
Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/
package main

import (
	"fmt"
)

//O(n) solution
func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func containsDuplicate1(nums []int) bool {
	mymap := make(map[int]int)
	for _, v := range nums {
		mymap[v]++

		if mymap[v] > 1 {
			return true
		}
	}
	fmt.Println(mymap)
	return false
}
func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))

}
