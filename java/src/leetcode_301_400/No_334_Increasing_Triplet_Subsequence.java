package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/12
 */
public class No_334_Increasing_Triplet_Subsequence {
    
    public boolean increasingTriplet(int[] nums) {
        boolean hasSecond = false;
        int len = nums.length;
        if (len < 3) { return false; }
        int first = nums[0], second = Integer.MAX_VALUE;
        for (int i = 1; i < len; i++) {
            if (!hasSecond) {
                if (nums[i] > first) { second = nums[i]; hasSecond = true; } 
                else if (nums[i] < first) { first = nums[i]; }
            } else {
                if (nums[i] > second) { return true; }
                else if (nums[i] > first) { second = nums[i]; }
                else { first = nums[i]; }
            }
        }
        return false;
    }
    
    public static void main(String[] args) {
        No_334_Increasing_Triplet_Subsequence solution = new No_334_Increasing_Triplet_Subsequence();
        int[][] inputs = new int[][]{
                new int[]{1, 2, 3, 4, 5},
                new int[]{5, 4, 3, 2, 1},
                new int[]{5, 1, 5, 5, 2, 5, 4}
        };
        for (int[] input : inputs) {
            System.out.println(solution.increasingTriplet(input));
        }
    }
    
}
