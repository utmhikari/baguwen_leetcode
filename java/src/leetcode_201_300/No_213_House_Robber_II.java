package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/6
 */
public class No_213_House_Robber_II {

    private int rob(int[] nums, int left, int right) {
        int inc = 0, exc = 0;
        for (int i = left; i <= right; i++) {
            int tmp = exc;
            exc = Math.max(inc, exc);
            inc = tmp + nums[i];
        }
        return Math.max(inc, exc);
    }

    public int rob(int[] nums) {
        int l = nums.length;
        if (l == 0) { return 0; }
        else if (l == 1) { return nums[0]; }
        else { return Math.max(rob(nums, 0, l - 2), rob(nums, 1, l - 1)); }
    }

    public static void main(String[] args) {
        No_213_House_Robber_II solution = new No_213_House_Robber_II();
        int[][] inputs = new int[][] {
                new int[] {2, 3, 2},
                new int[] {1, 2, 3, 1}
        };
        for (int[] i : inputs) {
            System.out.println(solution.rob(i));
        }
    }
}
