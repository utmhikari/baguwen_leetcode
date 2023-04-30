package leetcode_101_200;

import java.util.*;
/**
 * Created by 36191 on 2018/2/12
 * 123.最佳出售收益III: 2次
 */

public class No_123_Best_Time_to_Buy_and_Sell_Stock_III {

    public int maxProfit(int[] prices, int l, int r) {
        int max = 0;
        int sell = r;
        int point = r - 1;
        while(point >= l) {
            if(prices[point] < prices[sell]) {
                max = Math.max(max, prices[sell] - prices[point]);
            } else {
                sell = point;
            }
            point--;
        }
        return max;
    }

    public int maxProfit(int[] prices) {
        ArrayList<Integer> stops = new ArrayList<>();
        if(prices.length <= 1) {
            return 0;
        }
        int r = prices.length - 1;
        int l = prices.length - 1;
        while(true) {
            l--;
            if(l < 0) {
                break;
            }
            if(prices[l] > prices[l + 1]) {
                if(l < r - 1) {
                    stops.add(r);
                }
                r = l;
            }
        }
        if(l < r - 1) {
            stops.add(r);
        }
        int max = 0;
        if(!stops.isEmpty()) {
            for(int s : stops) {
                int profit_1 = maxProfit(prices, 0, s);
                int profit_2 = maxProfit(prices, s + 1, prices.length - 1);
                int profit = profit_1 + profit_2;
                if(profit > max) {
                    max = profit;
                }
            }
        }
        return max;
    }

    public static void main(String[] args) {
        No_123_Best_Time_to_Buy_and_Sell_Stock_III solution = new No_123_Best_Time_to_Buy_and_Sell_Stock_III();
        System.out.println(solution.maxProfit(new int[]{3, 3, 5, 0, 0, 3, 1, 4}));
        System.out.println(solution.maxProfit(new int[]{1, 2}));
        System.out.println(solution.maxProfit(new int[]{6, 1, 3, 2, 4, 7}));
        System.out.println(solution.maxProfit(new int[]{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}));
        System.out.println(solution.maxProfit(new int[]{2, 1}));
        System.out.println(solution.maxProfit(new int[]{}));
    }

}
