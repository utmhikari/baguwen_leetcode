package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/7
 */
public class No_220_Contains_Duplicate_III {
    public boolean containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        // t: diff of nums[i] & nums[j], k: diff of i & j
        int len = nums.length;
        if (len > 1) {
            for (int i = 0; i < len - 1; i++) {
                for (int j = i + 1; j <= Math.min(i + k, len - 1); j++) {
                    if (Math.abs((long)nums[i] - (long)nums[j]) <= (long)t) { return true; }
                }
            }
        }
        return false;
    }

    public static void main(String[] args) {
        No_220_Contains_Duplicate_III solution = new No_220_Contains_Duplicate_III();
        System.out.println(solution.containsNearbyAlmostDuplicate(
                new int[] {1, 5, 9, 1, 5, 9}, 2, 3
        ));

    }
}
