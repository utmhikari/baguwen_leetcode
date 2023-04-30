package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/27
 */

public class No_192_Horse_Robber {
    public int rob(int[] nums) {
        int len = nums.length;
        if (len == 0) { return 0; }
        if (len == 1) { return nums[0]; }
        if (len == 2) { return Math.max(nums[0], nums[1]); }
        int[] maxs = new int[len];
        maxs[0] = nums[0]; maxs[1] = nums[1]; maxs[2] = nums[0] + nums[2];
        for (int i = 3; i < len; i++) {
            maxs[i] = nums[i] + Math.max(maxs[i - 2], maxs[i - 3]);
        }
        return Math.max(maxs[len - 2], maxs[len - 1]);

    }
}
