package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/20
 */
public class No_238_Product_of_Array_Except_Self {
    public int[] productExceptSelf(int[] nums) {
        int n = 1;
        int[] ans = new int[nums.length];
        for (int i = nums.length - 1; i >= 0; i--) { ans[i] = n; n *= nums[i]; }
        n = 1;
        for (int i = 0; i < nums.length; i++) { ans[i] = ans[i] * n; n *= nums[i]; }
        return ans;
    }
}
