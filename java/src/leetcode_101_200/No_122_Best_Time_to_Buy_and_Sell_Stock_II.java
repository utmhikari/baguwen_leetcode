package leetcode_101_200;

/**
 * Created by 36191 on 2018/2/10
 * 122.最佳出售收益II：多次
 */

public class No_122_Best_Time_to_Buy_and_Sell_Stock_II {

    public int maxProfit(int[] prices) {
        if(prices.length <= 1) {
            return 0;
        }
        int sum = 0;
        int r = prices.length - 1;
        while(r > 0) {
            if(prices[r - 1] <= prices[r]) {
                sum += prices[r] - prices[r - 1];
            }
            r--;
        }
        return sum;
    }

    public static void main(String[] args) {
        No_122_Best_Time_to_Buy_and_Sell_Stock_II solution = new No_122_Best_Time_to_Buy_and_Sell_Stock_II();
        System.out.println(solution.maxProfit(new int[]{7, 1, 5, 3, 6, 4}));
        System.out.println(solution.maxProfit(new int[]{7, 6, 4, 3, 1}));
        System.out.println(solution.maxProfit(new int[]{2, 1, 4}));
        System.out.println(solution.maxProfit(new int[]{2, 4, 1}));
        System.out.println(solution.maxProfit(new int[]{2, 1, 2, 1, 0, 1, 2}));
        System.out.println(solution.maxProfit(new int[]{6, 1, 3, 2, 4, 7}));
        System.out.println(solution.maxProfit(new int[]{}));
    }
}
