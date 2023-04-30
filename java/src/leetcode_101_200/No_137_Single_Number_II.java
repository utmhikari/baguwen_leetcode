package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/17
 * No.137 唯一数字II，每个不唯一数字出现3次（2次太简单直接异或）
 */

public class No_137_Single_Number_II {

    public int singleNumber(int[] nums) {
        int x = 0, y = 0;
        for (int i : nums) {
            x = x ^ i & ~y;
            y = y ^ i & ~x;
        }
        return x | y;
    }

    public static void main(String[] args) {
        No_137_Single_Number_II solution = new No_137_Single_Number_II();
        System.out.println(solution.singleNumber(new int[] {0, 1, 0, 1, 99, 0, 1}));
    }

}
