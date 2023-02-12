# Given an m x n matrix mat, return an array of all the elements of the array in a diagonal order.
#
#  
# Example 1:
#
#
# Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
# Output: [1,2,4,7,5,3,6,8,9]
#
#
# Example 2:
#
#
# Input: mat = [[1,2],[3,4]]
# Output: [1,2,3,4]
#
#
#  
# Constraints:
#
#
# 	m == mat.length
# 	n == mat[i].length
# 	1 <= m, n <= 104
# 	1 <= m * n <= 104
# 	-105 <= mat[i][j] <= 105
#
#


class Solution:
    def findDiagonalOrder(self, matrix: List[List[int]]) -> List[int]:
        if matrix is None or matrix == []: return []
        res = []
        lines = defaultdict(list)
        rows = len(matrix)
        cols = len(matrix[0])
        for i in range(rows):
            for j in range(cols):
                lines[i+j].append(matrix[i][j])
        for k in range(rows + cols - 1):
            if k % 2 == 0:
                res += lines[k][::-1]
            else:
                res += lines[k]
        return res
