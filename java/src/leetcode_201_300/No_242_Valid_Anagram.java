package leetcode_201_300;

import java.util.ArrayList;

/**
 * Created by 36191 on 2018/10/20
 */

public class No_242_Valid_Anagram {
    public boolean isAnagram(String s, String t) {
        ArrayList<Character> chars = new ArrayList<>(), chart = new ArrayList<>();
        int[] cnts = new int[26], cntt = new int[26];
        int l = s.length();
        if (l != t.length()) { return false; }
        for (int i = 0; i < l; i++) {
            char cs = s.charAt(i), ct = t.charAt(i);
            chars.add(cs); chart.add(ct);
            cnts[cs - 'a'] += 1; cntt[ct - 'a'] += 1;
        }
        if (chars.size() != chart.size()) { return false; }
        for (char c : chars) {
            if (cnts[c - 'a'] != cntt[c - 'a']) { return false; }
        }
        return true;
    }
}
