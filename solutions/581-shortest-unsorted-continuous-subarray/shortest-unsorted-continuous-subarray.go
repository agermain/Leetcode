// Given an integer array nums, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order.
//
// Return the shortest such subarray and output its length.
//
//  
// Example 1:
//
//
// Input: nums = [2,6,4,8,10,9,15]
// Output: 5
// Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
//
//
// Example 2:
//
//
// Input: nums = [1,2,3,4]
// Output: 0
//
//
// Example 3:
//
//
// Input: nums = [1]
// Output: 0
//
//
//  
// Constraints:
//
//
// 	1 <= nums.length <= 104
// 	-105 <= nums[i] <= 105
//
//


import (
    "sort"
    "fmt"
)
func findUnsortedSubarray(nums []int) int {
    sorted := make([]int, len(nums))
    copy(sorted, nums)
    start, end := len(nums), 0
    sort.Ints(sorted)
    for idx, value := range sorted {
        if value != nums[idx] {
            start = min(start, idx)
            end = max(end, idx)
        }
    }
    if end-start >= 0 {
        return end-start+1
    } else {
        return 0
    }
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
