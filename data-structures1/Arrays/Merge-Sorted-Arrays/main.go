package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	var c []int
	var i, j = 0, 0

	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			c = append(c, nums1[i])
			i++
		} else if nums2[j] < nums1[i] {
			c = append(c, nums2[j])
			j++
		}
	}
	for ; i < m; i++ {
		c = append(c, nums1[i])
	}
	for ; j < n; j++ {
		c = append(c, nums2[j])
	}
	for i := 0; i < len(nums1); i++ {
		nums1[i] = c[i]
	}
}

//O(n) solution
func mergeSorted(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m == 0 || nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
	fmt.Println(nums1)
}

func main() {
	a := []int{2, 5, 8, 15, 18, 0, 0, 0, 0}
	b := []int{5, 9, 12, 17}
	mergeSorted(a, 5, b, 4)
}
