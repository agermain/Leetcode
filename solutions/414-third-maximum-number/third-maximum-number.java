// Given integer array nums, return the third maximum number in this array. If the third maximum does not exist, return the maximum number.
//
//  
// Example 1:
//
//
// Input: nums = [3,2,1]
// Output: 1
// Explanation: The third maximum is 1.
//
//
// Example 2:
//
//
// Input: nums = [1,2]
// Output: 2
// Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
//
//
// Example 3:
//
//
// Input: nums = [2,2,3,1]
// Output: 1
// Explanation: Note that the third maximum here means the third maximum distinct number.
// Both numbers with value 2 are both considered as second maximum.
//
//
//  
// Constraints:
//
//
// 	1 <= nums.length <= 104
// 	231 <= nums[i] <= 231 - 1
//
//
//  
// Follow up: Can you find an O(n) solution?


class Solution {
    private PriorityQueue<Integer> pq = new PriorityQueue<>();
    private Set<Integer> set = new HashSet<>();
    public int thirdMax(int[] nums) {
        for (int i : nums) {
            if (!set.contains(i)) {
                set.add(i);
                pq.add(i);
                if (pq.size() > 3) set.remove(pq.poll());
        }
    }
    if (pq.size() == 3) return pq.peek();
    while (pq.size() > 1) pq.poll();
    return pq.peek();
    }
}
