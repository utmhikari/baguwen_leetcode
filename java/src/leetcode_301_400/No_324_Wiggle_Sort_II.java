package leetcode_301_400;

import java.util.Arrays;

/**
 * Created by 36191 on 2019/1/12
 */
public class No_324_Wiggle_Sort_II {
    public void wiggleSort(int[] nums) {
        Arrays.sort(nums);
        int len = nums.length, mid = nums.length % 2 == 0 ? len / 2 - 1 : len / 2, idx = 0;
        int[] temp = new int[len];
        for (int i = 0; i <= mid; i++) {
            temp[idx] = nums[mid - i];
            if (idx + 1 < nums.length) {
                temp[idx + 1] = nums[len - i - 1];
            }
            idx += 2;
        }
        for (int i = 0; i < nums.length; i++) {
            nums[i] = temp[i];
        }
    }
}
