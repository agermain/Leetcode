// We are given two sentences A and B.  (A sentence is a string of space separated words.  Each word consists only of lowercase letters.)
//
// A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.
//
// Return a list of all uncommon words. 
//
// You may return the list in any order.
//
//  
//
//
//
//
//
// Example 1:
//
//
// Input: A = "this apple is sweet", B = "this apple is sour"
// Output: ["sweet","sour"]
//
//
//
// Example 2:
//
//
// Input: A = "apple apple", B = "banana"
// Output: ["banana"]
//
//
//  
//
// Note:
//
//
// 	0 <= A.length <= 200
// 	0 <= B.length <= 200
// 	A and B both contain only spaces and lowercase letters.
//
//
//
//


class Solution {
    public String[] uncommonFromSentences(String A, String B) {
        String[] words1 = A.split(" ");
        String[] word2 = B.split(" ");
        List<String> arrList = new ArrayList<String>();
        HashMap<String, Integer> hashMap = new HashMap<String, Integer>();
        
        for(String s:words1)
            hashMap.put(s,hashMap.getOrDefault(s,0)+1);
        for(String s:word2)
            hashMap.put(s,hashMap.getOrDefault(s,0)+1);
        for(String s: hashMap.keySet())
            if(hashMap.get(s)==1) arrList.add(s);
            
        return arrList.toArray(new String[0]);
    }
}
