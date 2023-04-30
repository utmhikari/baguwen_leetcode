package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/16
 * 053. 最大和的字数组
 */

public class No_053_Maximum_Subarray {

    public int maxSubArray(int[] nums) {
        int max = Integer.MIN_VALUE, sum = 0;
        for(int n : nums) {
            if(sum < 0) {
                sum = n;
            } else {
                sum += n;
            } if(sum > max) {
                max = sum;
            }
        }
        return max;
    }

    public static void main(String[] args) {
        No_053_Maximum_Subarray solution = new No_053_Maximum_Subarray();
        System.out.println(solution.maxSubArray(new int[] {-2,1,-3,4,-1,2,1,-5,4}));
    }
}
