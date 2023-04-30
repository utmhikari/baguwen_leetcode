package leetcode_301_400;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * Created by 36191 on 2019/1/18
 */
public class No_368_Largest_Divisible_Subset {
    
    public List<Integer> largestDivisibleSubset(int[] nums) {
        List<Integer> ans = new ArrayList<>();
        if (nums == null || nums.length == 0) { return ans; }
        Arrays.sort(nums);
        int[] dp = new int[nums.length];
        Arrays.fill(dp ,1);
        for (int i = 1; i < nums.length; i++){
            for (int j = i - 1; j >= 0; j--){
                if (nums[i] % nums[j] == 0){
                    dp[i] = Math.max(dp[i], dp[j] + 1);
                }
            }
        }
        int maxIndex = 0;
        for (int i = 1; i < nums.length; i++){
            maxIndex = dp[i] >= dp[maxIndex] ? i : maxIndex;
        }
        int temp = nums[maxIndex];
        int curDp = dp[maxIndex];
        for (int i = maxIndex; i >= 0; i--){
            if (temp % nums[i] == 0 && dp[i] == curDp){
                ans.add(nums[i]);
                temp = nums[i];
                curDp--;
            }
        }
        return ans;
    }
    
    public static void main(String[] args) {
        No_368_Largest_Divisible_Subset solution = new No_368_Largest_Divisible_Subset();
        int[][] inputs = new int[][] {
                new int[] {1, 2, 3},
                new int[] {8, 2, 4, 5, 6, 1},
        };
        for (int[] input : inputs) {
            List<Integer> ans = solution.largestDivisibleSubset(input);
            for (int i : ans) {
                System.out.print(i + ", ");
            }
            System.out.println();
        }
    }
}
