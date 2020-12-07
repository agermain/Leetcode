# Given an integer rowIndex, return the rowIndexth row of the Pascal's triangle.
#
# Notice that the row index starts from 0.
#
#
# In Pascal's triangle, each number is the sum of the two numbers directly above it.
#
# Follow up:
#
# Could you optimize your algorithm to use only O(k) extra space?
#
#  
# Example 1:
# Input: rowIndex = 3
# Output: [1,3,3,1]
# Example 2:
# Input: rowIndex = 0
# Output: [1]
# Example 3:
# Input: rowIndex = 1
# Output: [1,1]
#
#  
# Constraints:
#
#
# 	0 <= rowIndex <= 33
#
#


class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        row = [1] * (rowIndex+1)
        if rowIndex == 0: 
            return row
        previous = self.getRow(rowIndex-1)
        for i in range(1, len(row)-1):
            row[i] = previous[i-1] + previous[i]
        return row

        
