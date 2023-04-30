package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/2.
 * 001.数组中两个数相加得到目标值
 */
public class No_001_Two_Sum {
    public int[] twoSum(int[] nums, int target) {
        int[] ans = new int[2];
        for(int i = 0; i < nums.length - 1; i++) {
            for(int j = i + 1; j < nums.length; j++) {
                if((nums[i] + nums[j]) == target) {
                    ans[0] = i;
                    ans[1] = j;
                    return ans;
                }
            }
        }
        return ans;
    }

    public static void main(String[] args) {
        No_001_Two_Sum solution = new No_001_Two_Sum();
        int[] nums = {3, 2, 4};
        System.out.println(solution.twoSum(nums, 9));
    }
}
