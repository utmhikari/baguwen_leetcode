package leetcode_301_400;

/**
 * Created by 36191 on 2018/11/25
 */
public class No_309_Best_Time_to_Buy_and_Sell_Stock_with_Cooldown {

    public int maxProfit(int[] prices) {
        if (prices == null || prices.length == 0) {
            return 0;
        }
        int len = prices.length;
        int[] sell = new int[len], buy = new int[len];
        sell[0] = 0; buy[0] = -prices[0];
        for (int i = 1; i < prices.length; ++i) {
            sell[i] = Math.max(sell[i - 1], buy[i - 1] + prices[i]);
            buy[i] = Math.max(buy[i - 1], (i > 1 ? sell[i - 2] : 0) - prices[i]);
        }
        return sell[prices.length - 1];
    }

    public static void main(String[] args) {
        No_309_Best_Time_to_Buy_and_Sell_Stock_with_Cooldown solution =
                new No_309_Best_Time_to_Buy_and_Sell_Stock_with_Cooldown();
        int[][] inputs = new int[][] {
                new int[] {1, 2, 3, 0, 2},
                new int[] {6, 1, 3, 2, 4, 7},
        };
        for (int[] input : inputs) {
            System.out.println(solution.maxProfit(input));
        }
    }
}
