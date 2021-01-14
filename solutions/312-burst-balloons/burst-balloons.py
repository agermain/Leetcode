# You are given n balloons, indexed from 0 to n - 1. Each balloon is painted with a number on it represented by an array nums. You are asked to burst all the balloons.
#
# If you burst the ith balloon, you will get nums[i - 1] * nums[i] * nums[i + 1] coins. If i - 1 or i + 1 goes out of bounds of the array, then treat it as if there is a balloon with a 1 painted on it.
#
# Return the maximum coins you can collect by bursting the balloons wisely.
#
#  
# Example 1:
#
#
# Input: nums = [3,1,5,8]
# Output: 167
# Explanation:
# nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
# coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
#
# Example 2:
#
#
# Input: nums = [1,5]
# Output: 10
#
#
#  
# Constraints:
#
#
# 	n == nums.length
# 	1 <= n <= 500
# 	0 <= nums[i] <= 100
#
#


class Solution:
    def maxCoins(self, nums: List[int]) -> int:
        N = len(nums)
        fakeN = N+2
        if N <= 0:
            return 0
        # Add padding
        nums = [1] + nums + [1]
        table = [[0 for _ in range(fakeN)] for _ in range(fakeN)]

        # Windows (minus padding)
        for window in range(1, fakeN-1):
            for left in range(1, fakeN - window):
                right = left + window - 1 # Padding consideration
                for i in range(left,right+1):
                    current = table[left][i-1] + nums[left-1] * nums[i] * nums[right+1] + table[i+1][right]            
                    table[left][right] = max(table[left][right], current)
        return table[1][N]
