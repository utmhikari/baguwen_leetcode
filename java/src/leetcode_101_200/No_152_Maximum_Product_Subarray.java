package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/24
 * No.152 最大array乘积
 */

import java.util.*;

public class No_152_Maximum_Product_Subarray {

    public int maxProduct(int[] nums) {
        int max = nums[0];
        int[] dpMaxs = new int[nums.length], dpMins = new int[nums.length];
        dpMaxs[0] = nums[0]; dpMins[0] = nums[0];
        boolean flag = false;
        for(int i = 1; i < nums.length; i++) {
            if (nums[i] == 1) { dpMaxs[i] = dpMaxs[i - 1]; dpMins[i] = dpMins[i - 1]; flag = true;}
            else {
                int m1 = dpMaxs[i - 1] * nums[i], m2 = dpMins[i - 1] * nums[i];
                dpMaxs[i] = Math.max(nums[i], Math.max(m1, m2));
                dpMins[i] = Math.min(nums[i], Math.min(m1, m2));
                max = Math.max(dpMaxs[i], max);
            }
        }
        if (flag) { max = Math.max(1, max); }
        return max;
    }

    public static void main(String[] args) {
        No_152_Maximum_Product_Subarray solution = new No_152_Maximum_Product_Subarray();
        int[][] inputs = new int[][] {
                new int[] {2, 3, -2, 4},
                new int[] {-2, 0, -1},
                new int[] {1,2,-1,-2,2,1,-2,1,4,-5,4}
        };

        for (int[] input : inputs) {
            System.out.println(solution.maxProduct(input));
        }
    }
}
