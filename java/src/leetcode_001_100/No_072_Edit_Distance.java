package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/19
 * 072.字符串转换步数
 * 插入、删除、替换一个字母都算一步
 */

public class No_072_Edit_Distance {

    public int DPDistance(String shortS, String longS) {
        int sLen = shortS.length();
        int lLen = longS.length();
        if(sLen == 0) {
            return lLen;
        }
        int[][] dp = new int[lLen + 1][sLen + 1];
        dp[0][0] = 0;
        for(int i = 1; i < lLen + 1; i++) {
            dp[i][0] = i;
        }
        for(int j = 1; j < sLen + 1; j++) {
            dp[0][j] = j;
        }
        for(int j = 1; j < sLen + 1; j++) {
            for(int i = 1; i < lLen + 1; i++) {
                if(shortS.charAt(j - 1) == longS.charAt(i - 1)) {
                    dp[i][j] = dp[i - 1][j - 1];
                } else {
                    dp[i][j] = Math.min(Math.min(dp[i - 1][j - 1], dp[i - 1][j]), dp[i][j - 1]) + 1;
                }
            }
        }
        return dp[lLen][sLen];
    }

    public int minDistance(String word1, String word2) {
        if(word1.length() > word2.length()) {
            return DPDistance(word2, word1);
        } else {
            return DPDistance(word1, word2);
        }
    }

    public static void main(String[] args) {
        No_072_Edit_Distance solution = new No_072_Edit_Distance();
        System.out.println(solution.minDistance("abcde", "cbaecfd"));
        System.out.println(solution.minDistance("plasma", "altruism"));
    }
}
