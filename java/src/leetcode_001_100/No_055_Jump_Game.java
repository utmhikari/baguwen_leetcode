package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/16
 * 055. 能否跳跃到最后的index
 */

public class No_055_Jump_Game {

    public boolean canJump(int[] nums) {
        int len = nums.length;
        if(len == 0) {
            return false;
        } else if(len == 1) {
            return true;
        } else {
            int max = 0;
            for(int i = 0; i < len - 1; i++) {
                if(i > max) {
                    return false;
                }
                int temp = i + nums[i];
                if(temp > max) {
                    max = temp;
                    if(max >= len - 1) {
                        return true;
                    }
                }
            }
            return false;
        }
    }

    public static void main(String[] args) {
        No_055_Jump_Game solution = new No_055_Jump_Game();
        System.out.println(solution.canJump(new int[] {2, 3, 1, 1, 4} ));
        System.out.println(solution.canJump(new int[] {3, 2, 1, 0, 4} ));
        System.out.println(solution.canJump(new int[] {0} ));
        System.out.println(solution.canJump(new int[] {2,0,6,9,8,4,5,0,8,9,1,2,9,6,8,8,
                0,6,3,1,2,2,1,2,6,5,3,1,2,2,6,4,2,4,3,0,0,0,3,8,2,4,0,1,2,0,1,4,6,5,8,
                0,7,9,3,4,6,6,5,8,9,3,4,3,7,0,4,9,0,9,8,4,3,0,7,7,1,9,1,9,
                4,9,0,1,9,5,7,7,1,5,8,2,8,2,6,8,2,2,7,5,1,7,9,6} ));
        System.out.println(solution.canJump(new int[] {1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1
                ,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1
                ,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}));
    }
}
