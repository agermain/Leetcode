# You are given an m x n integer matrix matrix with the following two properties:
#
#
# 	Each row is sorted in non-decreasing order.
# 	The first integer of each row is greater than the last integer of the previous row.
#
#
# Given an integer target, return true if target is in matrix or false otherwise.
#
# You must write a solution in O(log(m * n)) time complexity.
#
#  
# Example 1:
#
#
# Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
# Output: true
#
#
# Example 2:
#
#
# Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
# Output: false
#
#
#  
# Constraints:
#
#
# 	m == matrix.length
# 	n == matrix[i].length
# 	1 <= m, n <= 100
# 	-104 <= matrix[i][j], target <= 104
#
#


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        if len(matrix) == 1 and target == 1:
            return True
        bs = []
        for row in matrix:
            bs += row
        
        low = 0
        high = len(bs) - 1
        mid = 0
        while low <= high:
            mid = (low + high)//2
            if bs[mid] == target:
                return True
            elif bs[mid] > target:
                high = mid-1
            else:
                low = mid+1
        return False
        
