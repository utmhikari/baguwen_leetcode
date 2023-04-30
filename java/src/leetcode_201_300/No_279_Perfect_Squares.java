package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/22
 */
public class No_279_Perfect_Squares {
    public int numSquares(int n) {
        int[] ans = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            int sqrt = (int)Math.sqrt((double)i);
            if (i % sqrt == 0 && i / sqrt == sqrt) { ans[i] = 1; }
            else {
                int min = Integer.MAX_VALUE;
                for (int j = 1; j <= sqrt; j++) { min = Math.min(min, ans[j * j] + ans[i - j * j]); }
                ans[i] = min;
            }
        }
        return ans[n];
    }
}
