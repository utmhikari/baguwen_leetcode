package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/24
 * No.153 在旋转后列表中寻找最小值
 */

public class No_153_Find_Minimum_in_Rotated_Sorted_Array {

    public int findMin(int[] nums) {
        for (int i = 0; i < nums.length - 1; i++) {
            if (nums[i] > nums[i + 1]) { return nums[i + 1]; }
        }
        return nums[0];
    }

    public static void main(String[] args) {
        No_153_Find_Minimum_in_Rotated_Sorted_Array solution = new No_153_Find_Minimum_in_Rotated_Sorted_Array();
        int[][] inputs = new int[][] {
                new int[] {3, 4, 5, 1, 2},
                new int[] {4, 5, 6, 7, 0, 1, 2}
        };

        for (int[] input : inputs) { System.out.println(solution.findMin(input)); }
    }

}
