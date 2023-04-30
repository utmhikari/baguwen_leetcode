package leetcode_401_500;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by 36191 on 2019/2/24
 */
public class No_438_Find_All_Anagrams_in_a_String {
    private boolean isAnagram(int[] chars) {
        for (int i = 0; i < 26; ++i) { if (chars[i] != 0) { return false; }}
        return true;
    }

    public List<Integer> findAnagrams(String s, String p) {
        List<Integer> ans = new ArrayList<>();
        int sl = s.length(), pl = p.length();
        if (sl < pl) { return ans; }
        int[] chars = new int[26];
        for (int i = 0; i < pl; ++i) {
            ++chars[p.charAt(i) - 'a'];
            --chars[s.charAt(i) - 'a'];
        }
        if (isAnagram(chars)) { ans.add(0); }
        for (int i = 0; i < sl - pl; ++i) {
            ++chars[s.charAt(i) - 'a'];
            --chars[s.charAt(i + pl) - 'a'];
            if (isAnagram(chars)) { ans.add(i + 1); }
        }
        return ans;
    }
}
