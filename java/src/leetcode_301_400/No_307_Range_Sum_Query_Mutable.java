package leetcode_301_400;

/**
 * Created by 36191 on 2018/11/25
 */
public class No_307_Range_Sum_Query_Mutable {

    private class NumArray {
        private int[] sums;

        public NumArray(int[] nums) {
            int l = nums.length, sum = 0;
            sums = new int[l];
            for (int i = 0; i < l; i++) {
                sum += nums[i];
                sums[i] = sum;
            }
        }

        public void update(int i, int val) {
            int prev = sumRange(i, i), l = sums.length;
            for (int x = i; x < l; x++) {
                sums[x] = sums[x] - prev + val;
            }
        }

        public int sumRange(int i, int j) {
            return i == 0 ? sums[j] : sumRange(0, j) - sumRange(0, i - 1);
        }
    }

    public NumArray newNumArray(int[] nums) {
        return new NumArray(nums);
    }

    public static void main(String[] args) {

    }
}
