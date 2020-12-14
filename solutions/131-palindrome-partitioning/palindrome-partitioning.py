# Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s.
#
# A palindrome string is a string that reads the same backward as forward.
#
#  
# Example 1:
# Input: s = "aab"
# Output: [["a","a","b"],["aa","b"]]
# Example 2:
# Input: s = "a"
# Output: [["a"]]
#
#  
# Constraints:
#
#
# 	1 <= s.length <= 16
# 	s contains only lowercase English letters.
#
#


class Solution:
    def partition(self, s: str) -> List[List[str]]:
        ans = []

        def dfs(currList, k):
            if k == len(s):
                ans.append(currList)
                return
            for i in range(k, len(s)):
                subStr = s[k:i+1]
                # Check if it's a palindrome by reversing it
                if subStr == subStr[::-1]:
                    dfs(currList+[subStr], i+1)

        dfs([], 0)
        return ans
