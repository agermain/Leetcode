# Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
#
# A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
#
#  
# Example 1:
# Input: s = "abc", t = "ahbgdc"
# Output: true
# Example 2:
# Input: s = "axc", t = "ahbgdc"
# Output: false
#
#  
# Constraints:
#
#
# 	0 <= s.length <= 100
# 	0 <= t.length <= 104
# 	s and t consist only of lowercase English letters.
#
#
#  
# Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109, and you want to check one by one to see if t has its subsequence. In this scenario, how would you change your code?


class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        m = []
        if len(s) == len(t):
            return s == t
        p1, p2 = 0, 0
        while p1 < len(s):
            if p2 >= len(t): return False
            char1, char2 = s[p1], t[p2]
            if char1 == char2:
                p1 += 1
                p2 += 1
                continue
            else:
                p2 += 1
                continue       

        return True

