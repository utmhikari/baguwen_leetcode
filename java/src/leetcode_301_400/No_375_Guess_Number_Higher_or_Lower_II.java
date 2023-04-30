package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/26
 */
public class No_375_Guess_Number_Higher_or_Lower_II {

    public int getMoneyAmount(int n) {
        int[][] dp = new int[n+1][n+1];
        return solve(dp, 1, n);
    }

    private int solve(int[][] dp, int left, int right){
        if (left >= right) { return 0; }
        if (dp[left][right] != 0) { return dp[left][right]; }
        int maxAmount = Integer.MAX_VALUE;
        for (int i = left; i <= right; i++) {
            maxAmount = Math.min(maxAmount, i + Math.max(solve(dp, left, i - 1), solve(dp, i + 1, right)));
        }
        dp[left][right] = maxAmount;
        return maxAmount;
    }
}
