package leetcode_301_400;

/**
 * Created by 36191 on 2019/2/7
 */
public class No_376_Wiggle_Subsequence {
    public int wiggleMaxLength(int[] nums) {
        int len = nums.length, d = 0, l = 1;
        if (len <= 1) { return len; }
        for (int i = 1; i < len; i++) {
            if (nums[i] > nums[i - 1] && d <= 0) { ++l; d = 1; }
            else if (nums[i] < nums[i - 1] && d >= 0) { ++l; d = -1; }
        }
        return l;
    }
}
