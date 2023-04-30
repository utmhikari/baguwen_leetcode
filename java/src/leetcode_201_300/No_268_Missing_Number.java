package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/21
 */
public class No_268_Missing_Number {
    public int missingNumber(int[] nums) {
        int sum = 0, all = nums.length * (nums.length + 1) / 2;
        for (int i : nums) { sum += i; }
        return all - sum;
    }
}
