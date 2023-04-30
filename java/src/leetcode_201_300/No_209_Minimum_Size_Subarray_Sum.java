package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/5
 */
public class No_209_Minimum_Size_Subarray_Sum {
    public int minSubArrayLen(int s, int[] nums) {
        int min = Integer.MAX_VALUE, left = 0, right = 0, len = nums.length, sum = 0;
        while (right < len) {
            sum += nums[right];
            if (sum >= s) {
                while (left <= right) {
                    sum -= nums[left++];
                    if (sum < s) { min = Math.min(min, right - left + 2); break; }
                    if (min == 1) { return 1; }
                }
            }
            ++right;
        }
        return min == Integer.MAX_VALUE ? 0 : min;
    }

    public static void main(String[] args) {
        No_209_Minimum_Size_Subarray_Sum solution = new No_209_Minimum_Size_Subarray_Sum();
        System.out.println(solution.minSubArrayLen(7, new int[]{2, 3, 1, 2, 4, 3}));
    }

}
