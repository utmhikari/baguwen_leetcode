package leetcode_401_500;

import java.util.Arrays;

/**
 * Created by 36191 on 2019/6/30
 */
public class No_416_Partition_Equal_Subset_Sum {


    private boolean find(int[] nums, int target, int current, int index) {
        if (index < 0) {
            return false;
        }
        for (int i = index; i >= 0; i--) {
            int sum = current + nums[i];
            if (sum == target) {
                return true;
            } else if (sum < target) {
                if (find(nums, target, sum, i - 1)) {
                    return true;
                }
            }
        }
        return false;
    }

    public boolean canPartition(int[] nums) {
        if (nums.length <= 1) {
            return false;
        }
        int sum = Arrays.stream(nums).sum();
        if (sum % 2 != 0 || nums[nums.length - 1] > sum / 2) {
            return false;
        }
        return find(nums, sum / 2, 0, nums.length - 1);
    }

    public static void main(String[] args) {
        No_416_Partition_Equal_Subset_Sum solution = new No_416_Partition_Equal_Subset_Sum();
        System.out.println(solution.canPartition(new int[]{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,100}));
    }
}
