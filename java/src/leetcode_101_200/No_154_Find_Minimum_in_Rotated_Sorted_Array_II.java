package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/24
 */

public class No_154_Find_Minimum_in_Rotated_Sorted_Array_II {
    public int findMin(int[] nums) {
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] < nums[i - 1]) {
                return nums[i];
            }
        }
        return nums[0];
    }
}
