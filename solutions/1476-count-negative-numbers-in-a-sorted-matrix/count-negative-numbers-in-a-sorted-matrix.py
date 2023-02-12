# Given a m x n matrix grid which is sorted in non-increasing order both row-wise and column-wise, return the number of negative numbers in grid.
#
#  
# Example 1:
#
#
# Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
# Output: 8
# Explanation: There are 8 negatives number in the matrix.
#
#
# Example 2:
#
#
# Input: grid = [[3,2],[1,0]]
# Output: 0
#
#
#  
# Constraints:
#
#
# 	m == grid.length
# 	n == grid[i].length
# 	1 <= m, n <= 100
# 	-100 <= grid[i][j] <= 100
#
#
#  
# Follow up: Could you find an O(n + m) solution?


class Solution:
    def countNegatives(self, grid: List[List[int]]) -> int:
        ans = 0
        print(len(grid))
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                if grid[i][j] < 0:
                    ans +=1
        return ans
