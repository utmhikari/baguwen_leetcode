package leetcode_001_100;

import java.util.Arrays;

/**
 * Created by 36191 on 2017/10/6.
 * 016.三个数和最接近目标
 */
public class No_016_3Sum_Closest {
    public int threeSumClosest(int[] nums, int target) {
        int ans;
        Arrays.sort(nums);
        int len = nums.length;
        if(len == 0) {
            return 0;
        }
        int minus = Integer.MAX_VALUE;
        int plus = Integer.MAX_VALUE;
        for(int i = 0 ; i < len - 2; i++) {
            if(i == 0 || nums[i] != nums[i-1]) {
                int front = i + 1, rear = len - 1;
                int sum = target - nums[i];
                while(front < rear) {
                    if(nums[front] + nums[rear] == sum) {
                        return target;
                    } else {
                        int diff = sum - nums[front] - nums[rear];
                        //System.out.println(diff);
                        if(diff > 0) {
                            if(diff < minus) {
                                minus = diff;
                            }
                            while(front < rear && nums[front] == nums[front + 1]) {
                                front++;
                            }
                            front++;
                        } else {
                            if(Math.abs(diff) < plus) {
                                plus = Math.abs(diff);
                            }
                            while(front < rear && nums[rear] == nums[rear - 1]) {
                                rear--;
                            }
                            rear--;
                        }
                    }
                }
            }
        }
        return (Math.abs(plus) < Math.abs(minus)) ? (target + plus) : (target - minus);
    }


    public static void main(String[] args) {
        No_016_3Sum_Closest solution = new No_016_3Sum_Closest();
        System.out.println(solution.threeSumClosest(new int[]{-1, 2, 1, -4}, 1));
        System.out.println(solution.threeSumClosest(new int[]{0, 0, 0}, 1));
        System.out.println(solution.threeSumClosest(new int[]{0, 2, 1, -3}, 1));
        System.out.println(solution.threeSumClosest(new int[]{0, 2, 1}, 0));
    }
}
