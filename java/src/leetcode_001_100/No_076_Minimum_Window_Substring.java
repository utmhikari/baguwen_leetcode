package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/21
 * 076. 包含所有字母最小窗口
 */

import java.util.HashMap;
public class No_076_Minimum_Window_Substring {

    /*
    SUBSTRING PROBLEMS:
    https://discuss.leetcode.com/topic/30941/here-is-a-10-line-template-that-can-solve-most-substring-problems
     */

    public String minWindow(String s, String t) {
        if(s.equals("") || t.equals("")) {
            return "";
        }
        HashMap<Character, Integer> charMap = new HashMap<>();
        int counter = t.length(), begin = 0, end = 0, d = Integer.MAX_VALUE, head = 0;
        for(char c = 33; c < 127; c++) {
            charMap.put(c, 0);
        }
        for(int i = 0; i < counter; i++) {
            charMap.put(t.charAt(i), charMap.get(t.charAt(i)) + 1); // 存储出现字母及其个数
        }
        while(end < s.length()){
            if(charMap.get(s.charAt(end)) > 0) {
                counter--;
            }
            charMap.put(s.charAt(end), charMap.get(s.charAt(end)) - 1);
            end++;
            while(counter == 0) {
                if(end - begin < d) {
                    d = end - begin;
                    head = begin;
                }
                if(charMap.get(s.charAt(begin)) == 0) {
                    counter++;
                }
                charMap.put(s.charAt(begin), charMap.get(s.charAt(begin)) + 1);
                begin++;
            }
        }
        return d == Integer.MAX_VALUE ? "" : s.substring(head, head + d);
    }

    public static void main(String[] args) {
        No_076_Minimum_Window_Substring solution = new No_076_Minimum_Window_Substring();
        System.out.println(solution.minWindow("ADOBECODEBANC", "ABC"));
        System.out.println(solution.minWindow("aa", "aa"));
    }
}
