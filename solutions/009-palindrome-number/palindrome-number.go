// Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
//
// Follow up: Could you solve it without converting the integer to a string?
//
//  
// Example 1:
//
//
// Input: x = 121
// Output: true
//
//
// Example 2:
//
//
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
//
//
// Example 3:
//
//
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//
//
// Example 4:
//
//
// Input: x = -101
// Output: false
//
//
//  
// Constraints:
//
//
// 	-231 <= x <= 231 - 1
//
//


func isPalindrome(x int) bool {
    if (x < 0 || (x % 10 == 0 && x != 0)) {
        return false
    } 
    var revertedNumber int
    for x > revertedNumber {
        revertedNumber = revertedNumber * 10 + x % 10
        x /= 10
    }
    return x == revertedNumber || x == revertedNumber/10
}
