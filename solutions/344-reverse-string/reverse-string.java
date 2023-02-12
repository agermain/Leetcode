// Write a function that reverses a string. The input string is given as an array of characters s.
//
// You must do this by modifying the input array in-place with O(1) extra memory.
//
//  
// Example 1:
// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]
// Example 2:
// Input: s = ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]
//
//  
// Constraints:
//
//
// 	1 <= s.length <= 105
// 	s[i] is a printable ascii character.
//
//


class Solution {
    public void reverseString(char[] s) {
        helper(0, s.length-1, s);
    }
    public void helper(int start, int end, char[] s) {
        if (start >= end) return;
        char tmp = s[end];
        s[end] = s[start];
        s[start] = tmp;
        helper(start+1, end-1, s);
    }
}
