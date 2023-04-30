package leetcode_301_400;

/**
 * Created by 36191 on 2018/11/17
 */
public class No_303_Range_Sum_Query_Immutable {
    class NumArray {
        private int[] sums;

        public NumArray(int[] nums) {
            int l = nums.length;
            if (l != 0) {
                sums = new int[l];
                sums[0] = nums[0];
                for (int i = 1; i < l; i++) { sums[i] = nums[i] + sums[i - 1]; }
            }
        }

        public int sumRange(int i, int j) {
            return i == 0 ? sums[j] : sums[j] - sums[i - 1];
        }
    }
}
