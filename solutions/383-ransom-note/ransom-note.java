// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
//
// Each letter in magazine can only be used once in ransomNote.
//
//  
// Example 1:
// Input: ransomNote = "a", magazine = "b"
// Output: false
// Example 2:
// Input: ransomNote = "aa", magazine = "ab"
// Output: false
// Example 3:
// Input: ransomNote = "aa", magazine = "aab"
// Output: true
//
//  
// Constraints:
//
//
// 	1 <= ransomNote.length, magazine.length <= 105
// 	ransomNote and magazine consist of lowercase English letters.
//
//


class Solution {
    public boolean canConstruct(String ransomNote, String magazine) {
        if ("".equals(ransomNote)) return true;
        int[] alphabet = new int[26];     
        for(int i=0; i<magazine.length(); i++){
            alphabet[magazine.charAt(i) - 'a']++;
        }
        for(int i=0; i<ransomNote.length(); i++){
            if(alphabet[ransomNote.charAt(i) - 'a'] == 0){
                return false;
            }
            alphabet[ransomNote.charAt(i) - 'a']--;
        }
        return true;
    }
}
