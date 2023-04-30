package leetcode_101_200;

import java.util.*;
/**
 * Created by 36191 on 2018/2/10
 * 121.最高出售收益I：一次
 */

public class No_121_Best_Time_to_Buy_and_Sell_Stock {

    public int maxProfit(int[] prices) {
        int max = 0;
        int sell = prices.length - 1;
        int point = prices.length - 2;
        while(point >= 0) {
            if(prices[point] < prices[sell]) {
                max = Math.max(max, prices[sell] - prices[point]);
            } else {
                sell = point;
            }
            point--;
        }
        return max;
    }

    public static void main(String[] args) {
        No_121_Best_Time_to_Buy_and_Sell_Stock solution = new No_121_Best_Time_to_Buy_and_Sell_Stock();
        System.out.println(solution.maxProfit(new int[]{7, 1, 5, 3, 6, 4}));
        System.out.println(solution.maxProfit(new int[]{7, 6, 4, 3, 1}));
        System.out.println(solution.maxProfit(new int[]{2, 1, 4}));
        System.out.println(solution.maxProfit(new int[]{2, 4, 1}));
        System.out.println(solution.maxProfit(new int[]{2, 1, 2, 1, 0, 1, 2}));
    }
}
