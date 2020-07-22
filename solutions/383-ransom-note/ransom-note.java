// Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.
//
// Each letter in the magazine string can only be used once in your ransom note.
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
// 	You may assume that both strings contain only lowercase letters.
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
