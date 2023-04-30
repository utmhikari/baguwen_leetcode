package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/14
 * 045. 最小跳跃次数
 */

public class No_045_Jump_Game_II {


    public int jump(int[] nums) {
        int len = nums.length;
        if(len <= 1) {
            return 0;
        }
        int count = 0;
        int target = 0;
        int max = 0;
        for(int i = 0; i < len - 1; i++) {
            max = Math.max(i + nums[i], max);
            if(i == target) {
                count++;
                target = max;
            }
        }
        return count;
    }

    public static void main(String[] args) {
        No_045_Jump_Game_II solution = new No_045_Jump_Game_II();
        System.out.println(solution.jump(new int[] {2, 3, 1, 1, 4}));
        System.out.println(solution.jump(new int[] {1, 3, 2, 2, 4, 5, 3, 2, 2}));
    }
}
