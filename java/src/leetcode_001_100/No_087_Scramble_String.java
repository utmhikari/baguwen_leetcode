package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/28
 * 087. 二叉树分解字符串（转换left right）
 */

public class No_087_Scramble_String {


    public boolean isScramble(String s1, String s2) {
        if (s1.equals(s2)) {
            return true;
        }
        int[] cnt = new int[26];
        for (char c : s1.toCharArray()) {
            cnt[c - 'a']++;
        }
        for (char c : s2.toCharArray()) {
            cnt[c - 'a']--;
        }
        for (int i : cnt) {
            if (i != 0) {
                return false;
            }
        }
        int len = s1.length(), mid = len / 2;
        for (int p = 1; p <= mid; ++p) {
            String s1_l1 = s1.substring(0, p), s1_r1 = s1.substring(p),
                    s1_l2 = s1.substring(0, len - p), s1_r2 = s1.substring(len - p),
                    s2_l1 = s2.substring(0, p), s2_r1 = s2.substring(p),
                    s2_l2 = s2.substring(0, len - p), s2_r2 = s2.substring(len - p);
            if ((isScramble(s1_l1, s2_l1) && isScramble(s1_r1, s2_r1)) ||
                    (isScramble(s1_r2, s2_l1) && isScramble(s1_l2, s2_r1)) ||
                    (isScramble(s1_l1, s2_r2) && isScramble(s1_r1, s2_l2)) ||
                    (isScramble(s1_l2, s2_l2) && isScramble(s1_r2, s2_r2))) {
                return true;
            }
        }
        return false;
    }

    public static void main(String[] args) {
        No_087_Scramble_String solution = new No_087_Scramble_String();
        System.out.println(solution.isScramble("great", "rgeat"));
        System.out.println(solution.isScramble("great", "rgtae"));
        System.out.println(solution.isScramble("great", "taerg"));
        System.out.println(solution.isScramble("abs", "bas"));
        System.out.println(solution.isScramble("abc", "bca"));
        System.out.println(solution.isScramble("abb", "bab"));
    }
}
