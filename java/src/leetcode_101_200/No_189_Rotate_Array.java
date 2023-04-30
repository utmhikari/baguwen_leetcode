package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/26
 */

public class No_189_Rotate_Array {
    public void rotate(int[] nums, int k) {
        int l = nums.length;
        int[] rot = new int[l];
        for (int i = 0; i < l; i++) { rot[(i + k) % l] = nums[i]; }
        System.arraycopy(rot, 0, nums, 0, l);
    }
}
