package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/4.
 * 010.正则表达式匹配——动态规划
 */

public class No_010_Regular_Expression_Matching {

    public boolean isMatch(String s, String p) {
        if (s == null || p == null) {
            return false;
        }
        boolean[][] dp = new boolean[s.length()+1][p.length()+1]; // 初始化所有匹配状况为false
        dp[0][0] = true; // 空串匹配空串
        for (int i = 0; i < p.length(); i++) {
            if (p.charAt(i) == '*' && dp[0][i-1]) {
                dp[0][i+1] = true;
            }
        }
        for (int i = 0 ; i < s.length(); i++) {
            for (int j = 0; j < p.length(); j++) {
                // 等于'.'或该字符时，下一行下一列的bool相应更新
                if (p.charAt(j) == '.') {
                    dp[i+1][j+1] = dp[i][j];
                    continue;
                }
                if (p.charAt(j) == s.charAt(i)) {
                    dp[i+1][j+1] = dp[i][j];
                    continue;
                }
                // 等于'*'时分情况讨论
                if (p.charAt(j) == '*') {
                    if (p.charAt(j-1) != s.charAt(i) && p.charAt(j-1) != '.') {
                        // *前的字符不匹配，将其看作空，并将之前的匹配结果传入
                        dp[i+1][j+1] = dp[i+1][j-1];
                    } else {
                        // *前的字符匹配，可以看作单个字符或多个字符或空
                        dp[i+1][j+1] = (dp[i+1][j] || dp[i][j+1] || dp[i+1][j-1]);
                    }
                }
            }
        }
        return dp[s.length()][p.length()];
    }

    public static void main(String[] args) {
        No_010_Regular_Expression_Matching solution = new No_010_Regular_Expression_Matching();
        System.out.println(solution.isMatch("aa","a"));
        System.out.println(solution.isMatch("aa","aa"));
        System.out.println(solution.isMatch("aaa","aa"));
        System.out.println(solution.isMatch("aa","a*"));
        System.out.println(solution.isMatch("aa",".*"));
        System.out.println(solution.isMatch("ab",".*"));
        System.out.println(solution.isMatch("c","c*a*b"));
    }
}
