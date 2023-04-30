package leetcode_301_400;

/**
 * Created by 36191 on 2019/2/7
 */
public class No_392_is_subsequence {
    public boolean isSubsequence(String s, String t) {
        int lt = t.length(), ls = s.length(), idx = 0;
        if (ls == 0) { return true; }
        for (int i = 0; i < lt; i++) {
            if (s.charAt(idx) == t.charAt(i)) { ++idx; }
            if (idx >= ls) { return true; }
        }
        return false;
    }
}
