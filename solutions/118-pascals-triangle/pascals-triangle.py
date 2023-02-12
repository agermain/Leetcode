# Given an integer numRows, return the first numRows of Pascal's triangle.
#
# In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
#
#  
# Example 1:
# Input: numRows = 5
# Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
# Example 2:
# Input: numRows = 1
# Output: [[1]]
#
#  
# Constraints:
#
#
# 	1 <= numRows <= 30
#
#


class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        if numRows == 1:
            return [[1]]
        ans = [[1]]
        for row in range(2, numRows+1):
            current = []
            for num in range(row):
                print(num)
                if num == 0 or num == row-1:
                    current.append(1)
                else:
                    number = ans[row-2][num-1] + ans[row-2][num]
                    current.append(number)
            ans.append(current)
        return ans
