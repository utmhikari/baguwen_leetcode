package leetcode_301_400;

import java.util.Arrays;
import java.util.HashMap;

/**
 * Created by 36191 on 2019/2/8
 */
public class No_377_Combination_Sum_IV {
    private int dfs(int[] nums, int target, int idx, int sum, HashMap<Integer, Integer> cnts) {
        if (cnts.containsKey(sum)) {
            return cnts.get(sum);
        } else {
            int count = 0;
            for (int i = 0; i < nums.length; i++) {
                int tmp = sum + nums[i];
                if (tmp < target) {
                    count += dfs(nums, target, i, tmp, cnts);
                }
                else {
                    if (tmp == target) { count++; }
                    break;
                }
            }
            cnts.put(sum, count);
            return count;
        }
    }

    public int combinationSum4(int[] nums, int target) {
        if (target == 0) { return 0; }
        Arrays.sort(nums);
        HashMap<Integer, Integer> cnts = new HashMap<>();
        return dfs(nums, target, 0, 0, cnts);
    }
}
