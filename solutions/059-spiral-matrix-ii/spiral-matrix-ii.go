// Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.
//
//  
// Example 1:
//
//
// Input: n = 3
// Output: [[1,2,3],[8,9,4],[7,6,5]]
//
//
// Example 2:
//
//
// Input: n = 1
// Output: [[1]]
//
//
//  
// Constraints:
//
//
// 	1 <= n <= 20
//
//


func generateMatrix(n int) [][]int {
    res := make([][]int, n)
    for i := 0; i < n; i++ {
        res[i] = make([]int, n)
    }

    if n == 0 {
        return res
    }
    top, bottom, left, right, current := 0, n-1, 0, n-1, 1
    for left <= right && top <= bottom {
        // Shrink columns and rows
        
        // First row, from left to right
        for j:=left; j <= right; j++ {
            res[top][j] = current
            current++
        }
        top++
        // Most right column, from top to bottom
        for i:=top; i <= bottom; i++ {
            res[i][right] = current
            current++
        }
        right--
        
        // Last row, from right to left
        for j:=right; j >= left; j-- {
            res[bottom][j] = current
            current++
        }
        bottom--
        
        // Most left column, from bottom to top
        for i:=bottom; i >= top ; i-- {
            res[i][left] = current
            current++
        }
        left++
        
    }
    return res
}
