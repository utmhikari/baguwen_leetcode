package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/3.
 * 003.最长不重复字符的子串
 */

import java.util.HashSet;
public class No_003_Longest_Unrepeated_Substr {
    public int lengthOfLongestSubstring(String s) {
        if(s.length() == 0) {
            return 0;
        }
        if(s.length() == 1) {
            return 1;
        }
        int maxLength = 0;
        HashSet<Character> charSet = new HashSet<>();
        int strLength = s.length();
        for(int i = 0; i < strLength; i++) {
            for(int j = 0; j < strLength - i; j++) {
                charSet.add(s.charAt(i + j));
                int substrLength = charSet.size();
                if(substrLength != j + 1) {
                    if(substrLength > maxLength) {
                        maxLength = substrLength;
                    }
                    charSet.clear();
                    charSet = new HashSet<>();
                    break;
                }
            }
            if(charSet.size() > maxLength) {
                maxLength = charSet.size();
                charSet.clear();
                charSet = new HashSet<>();
            }
            if(i + maxLength >= strLength - 1) {
                break;
            }
        }
        return maxLength;
    }

    public static void main(String[] args) {
        No_003_Longest_Unrepeated_Substr solution = new No_003_Longest_Unrepeated_Substr();
        String a = "abcabcbb";
        String b = "bbbbb";
        String c = "pwwkew";
        String d = "aab";
        System.out.println(solution.lengthOfLongestSubstring(a));
        System.out.println(solution.lengthOfLongestSubstring(b));
        System.out.println(solution.lengthOfLongestSubstring(c));
        System.out.println(solution.lengthOfLongestSubstring(d));
    }
}
