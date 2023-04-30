package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/9.
 * 030.包含所有单词子串
 */

import java.util.List;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
public class No_030_Substr_Concat_All_Words {

    public static List<Integer> findSubstring(String S, String[] L) {
        List<Integer> res = new ArrayList<Integer>();
        if (S == null || L == null || L.length == 0) return res;
        int len = L[0].length(); // length of each word

        Map<String, Integer> map = new HashMap<>(); // map for L
        for (String w : L) map.put(w, map.containsKey(w) ? map.get(w) + 1 : 1);

        for (int i = 0; i <= S.length() - len * L.length; i++) {
            Map<String, Integer> copy = new HashMap<>(map);
            for (int j = 0; j < L.length; j++) { // checkc if match
                String str = S.substring(i + j*len, i + j*len + len); // next word
                if (copy.containsKey(str)) { // is in remaining words
                    int count = copy.get(str);
                    if (count == 1) copy.remove(str);
                    else copy.put(str, count - 1);
                    if (copy.isEmpty()) { // matches
                        res.add(i);
                        break;
                    }
                } else break; // not in L
            }
        }
        return res;
    }

    /*
    public boolean isMatch(String s, String[] words, int startIndex, HashMap<String, Integer> wordMap) {
        int len = words[0].length();
        for(int j = startIndex; j < words.length; j++) {
            String subStr = s.substring(startIndex + j * len, startIndex + j * len + len);
            if(wordMap.containsKey(subStr)) {
                int value = wordMap.get(subStr);
                if(value == 1) {
                    wordMap.remove(subStr);
                } else {
                    wordMap.put(subStr, value - 1);
                }
            } else {
                return false;
            }
        }
        if(wordMap.isEmpty()) {
            return true;
        }
        return false;
    }

    public List<Integer> findSubstring(String s, String[] words) {
        HashMap<String, Integer> wordMap = new HashMap<>();
        for(String w : words) {
            if(!wordMap.containsKey(w)) {
                wordMap.put(w, 1);
            } else {
                int newValue = wordMap.get(w) + 1;
                wordMap.put(w, newValue);
            }
        }
        List<Integer> fs = new ArrayList<>();
        if(words.length == 0) {
            return fs;
        }
        int len = words.length;
        int lenW = words[0].length();
        for(int i = 0; i < s.length() - len * lenW + 1; i++) {
            if(isMatch(s, words, i, new HashMap<>(wordMap))) {
                fs.add(i);
            }
        }
        return fs;
    }
    */

    public void test(String s, String[] words) {
        List<Integer> fs = findSubstring(s, words);
        for(int i : fs) {
            System.out.print(i + ", ");
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_030_Substr_Concat_All_Words solution = new No_030_Substr_Concat_All_Words();
        solution.test("barfoothefoobarman", new String[] {"foo", "bar"});
        solution.test("ababaab", new String[] {"ab", "ba", "ba"});
    }
}
