# Given a positive integer n, find the smallest integer which has exactly the same digits existing in the integer n and is greater in value than n. If no such positive integer exists, return -1.
#
# Note that the returned integer should fit in 32-bit integer, if there is a valid answer but it does not fit in 32-bit integer, return -1.
#
#  
# Example 1:
# Input: n = 12
# Output: 21
# Example 2:
# Input: n = 21
# Output: -1
#
#  
# Constraints:
#
#
# 	1 <= n <= 231 - 1
#
#


class Solution:
    def nextGreaterElement(self, n: int) -> int:
        if n == 0: return -1
        nums = list(str(n))
        ln = len(nums)
        i = ln-1
        while i > 0:
            if nums[i-1] < nums[i]: break
            i -= 1
        i -= 1  
        if i < 0: return -1
        tmp = ln-1
        while tmp > i:
            if nums[i] < nums[tmp]: break
            tmp -= 1
        nums[i], nums[tmp] = nums[tmp], nums[i]
        nums[i+1:] = sorted(nums[i+1:]) 
        res = int("".join(nums))
        return res if (res > n and res <= (2**31-1)) else -1
