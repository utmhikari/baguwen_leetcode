package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/7.
 * 018.4个数和为0
 */
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
public class No_018_4Sum {

    public List<List<Integer>> fourSum(int[] nums, int target) {
        List<List<Integer>> ans = new ArrayList<>();
        Arrays.sort(nums);
        int len = nums.length;
        if(len < 4 || nums[0] > 0) {
            return ans;
        }
        for(int i0 = 0; i0 < len - 3; i0++) {
            for(int i1 = i0 + 1 ; i1 < len - 2; i1++) {
                if(i1 == i0 + 1 || nums[i1] != nums[i1-1]) {
                    int front = i1 + 1, rear = len - 1;
                    int sum = target - nums[i0] - nums[i1];
                    while(front < rear) {
                        if(nums[front] + nums[rear] == sum) {
                            ans.add(Arrays.asList(nums[i0], nums[i1], nums[front], nums[rear]));
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
                while(i1 < len - 2 && nums[i1] == nums[i1 + 1]) {
                    i1++;
                }
            }
            while(i0 < len - 3 && nums[i0] == nums[i0 + 1]) {
                i0++;
            }
        }
        return ans;
    }

    public void output(int[] nums, int target) {
        List<List<Integer>> ans = fourSum(nums, target);
        for(List<Integer> l : ans) {
            for(int i : l) {
                System.out.print(String.valueOf(i) + ", ");
            }
            System.out.print("\n");
        }
    }

    public static void main(String[] args) {
        No_018_4Sum solution = new No_018_4Sum();
        solution.output(new int[]{-1, 0, 1, 0, -2, 2}, 0);
        solution.output(new int[]{0, 0, 0, 0}, 0);
    }
}
