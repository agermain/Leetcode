// Given two strings s and goal, return true if and only if s can become goal after some number of shifts on s.
//
// A shift on s consists of moving the leftmost character of s to the rightmost position.
//
//
// 	For example, if s = "abcde", then it will be "bcdea" after one shift.
//
//
//  
// Example 1:
// Input: s = "abcde", goal = "cdeab"
// Output: true
// Example 2:
// Input: s = "abcde", goal = "abced"
// Output: false
//
//  
// Constraints:
//
//
// 	1 <= s.length, goal.length <= 100
// 	s and goal consist of lowercase English letters.
//
//


class Solution {
    public boolean rotateString(String A, String B) {
        StringBuilder sb = new StringBuilder(A);
        int len = 0;
        if (A.length() == 0&& B.length() == 0) return true;
        for (int i=0; i<sb.length(); i++) {
            char c = sb.charAt(0);
            sb.deleteCharAt(0);
            sb.append(c);
            if (sb.toString().equals(B)) return true;
        }
        return false;
    }
}
