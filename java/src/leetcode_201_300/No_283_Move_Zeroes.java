package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/22
 */
public class No_283_Move_Zeroes {
    public void moveZeroes(int[] nums) {
        int cur = 0, l = nums.length, prev = 0;
        while (true) {
            while (cur < l && nums[cur] == 0) { cur++; }
            if (cur == nums.length) { break; }
            if (cur == prev) { ++prev; ++cur; }
            else { nums[prev++] = nums[cur]; nums[cur++] = 0; }
        }
    }
}
