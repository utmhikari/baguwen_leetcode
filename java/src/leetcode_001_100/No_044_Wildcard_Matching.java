package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/14
 * 044. 通配符匹配
 */

public class No_044_Wildcard_Matching {
    public boolean isMatch(String s, String p) {
        if (s == null || p == null) {
            return false;
        }
        int sLen = s.length();
        int pLen = p.length();
        boolean[][] dp = new boolean[sLen + 1][pLen + 1];
        dp[0][0] = true;
        for(int j = 1; j < pLen + 1; j++) {
            if(p.charAt(j - 1) == '*') {
                if(j == 1) {
                    dp[0][j] = true;
                } else {
                    dp[0][j] = dp[0][j - 1];
                }
            } else {
                dp[0][j] = false;
            }
        }
        for(int i = 1; i < sLen + 1; i++) {
            dp[i][0] = false;
        }
        for(int i = 1; i < sLen + 1; i++) {
            for(int j = 1; j < pLen + 1; j++) {
                if(p.charAt(j - 1) == '?') {
                    dp[i][j] = dp[i - 1][j - 1];
                } else if(p.charAt(j - 1) == '*') {
                    dp[i][j] = dp[i][j - 1] || dp[i - 1][j];
                } else {
                    dp[i][j] = dp[i - 1][j - 1] && (s.charAt(i - 1) == p.charAt(j - 1));
                }
            }
        }
        return dp[sLen][pLen];
    }

    public static void main(String[] args) {
        No_044_Wildcard_Matching solution = new No_044_Wildcard_Matching();
        System.out.println(solution.isMatch("aa","a"));
        System.out.println(solution.isMatch("aa","aa"));
        System.out.println(solution.isMatch("aaa","aa"));
        System.out.println(solution.isMatch("aa", "*"));
        System.out.println(solution.isMatch("aa", "a*"));
        System.out.println(solution.isMatch("ab", "?*"));
        System.out.println(solution.isMatch("aab", "c*a*b"));
    }
}
