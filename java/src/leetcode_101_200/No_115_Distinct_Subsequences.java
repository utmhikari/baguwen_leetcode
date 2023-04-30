package leetcode_101_200;

/**
 * Created by 36191 on 2017/12/13
 * 115.不同的子序列数
 */

public class No_115_Distinct_Subsequences {

    public int numDistinct(String s, String t) {
        if(s.length() == 0 || t.length() == 0) {
            return 0;
        }
        int[][] dp = new int[t.length() + 1][s.length() + 1];
        for(int r = 1; r <= t.length(); r++) {
            for(int c = r; c <= s.length(); c++) {
                if(s.charAt(c - 1) != t.charAt(r - 1)) {
                    dp[r][c] = dp[r][c - 1];
                } else {
                    if(r == 1) {
                        dp[r][c] = dp[r][c - 1] + 1;
                    } else {
                        dp[r][c] = dp[r][c - 1] + dp[r - 1][c - 1];
                    }
                }
            }
        }
        return dp[t.length()][s.length()];
    }

    public static void main(String[] args) {
        No_115_Distinct_Subsequences solution = new No_115_Distinct_Subsequences();
        System.out.println(solution.numDistinct("rabbbit", "rabbit"));
        System.out.println(solution.numDistinct("rabbbiit", "rabbit"));
        System.out.println(solution.numDistinct("rarabbbit", "rabbit"));

    }
}
