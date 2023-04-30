package leetcode_401_500;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by 36191 on 2019/2/24
 */
public class No_448_Find_All_Numbers_Disappeared_in_an_Array {
    private void swap(int[] nums, int i1, int i2) {
        int temp = nums[i1];
        nums[i1] = nums[i2];
        nums[i2] = temp;
    }

    public List<Integer> findDisappearedNumbers(int[] nums) {
        List<Integer> ans = new ArrayList<>();
        int i = 0, l = nums.length;
        // traverse 2 times
        // first time put each digit to its position (idx + 1), tag 0 on those disappeareds
        // second time add those disappeareds with 0 on num[idx + 1]
        // final result of nums will be like: [1, 2, 3, 4, 0, 0, 7, 8] so that we can find 5, 6
        while (i < l) {
            if (nums[i] == 0 || nums[i] == i + 1) {
                // if meet a 0 or the digit is at its right position, go forward!
                ++i;
            } else if (nums[nums[i] - 1] == nums[i]) {
                // if the digit is already at its position, fuck the index to 0
                nums[i++] = 0;
            } else {
                // swap the digit to its right position
                swap(nums, i, nums[i] - 1);
            }
        }
        for (i = 0; i < l; ++i) { if (nums[i] == 0) { ans.add(i + 1); }}
        return ans;
    }
}
