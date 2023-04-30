package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/6.
 * 015.三个数和为0
 */

import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

public class No_015_3Sum {
    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> ans = new ArrayList<>();
        Arrays.sort(nums);
        int len = nums.length;
        if(len == 0) {
            return ans;
        }
        if(nums[0] > 0) {
            return ans;
        } else {
            for(int i = 0 ; i < len - 2; i++) {
                if(i == 0 || nums[i] != nums[i-1]) {
                    int front = i + 1, rear = len - 1;
                    int sum = 0 - nums[i];
                    while(front < rear) {
                        if(nums[front] + nums[rear] == sum) {
                            ans.add(Arrays.asList(nums[i], nums[front], nums[rear]));
                            while(front < rear && nums[front] == nums[front + 1]) {
                                front++;
                            }
                            while(front < rear && nums[rear] == nums[rear - 1]) {
                                rear--;
                            }
                            front++;
                            rear--;
                        } else {
                            if(nums[front] + nums[rear] < sum) {
                                front++;
                            } else {
                                rear--;
                            }
                        }
                    }
                }
            }
        }
        return ans;
    }

    public void output(int[] nums) {
        List<List<Integer>> ans = threeSum(nums);
        for(List<Integer> l : ans) {
            for(int i : l) {
                System.out.print(String.valueOf(i) + ", ");
            }
            System.out.print("\n");
        }
    }

    public static void main(String[] args) {
        No_015_3Sum solution = new No_015_3Sum();
        solution.output(new int[]{-1, 0, 2, 1, -1, -4});

    }
}
