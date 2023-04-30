package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/10.
 * 031.下一个排列
 */

import java.util.Arrays;
public class No_031_Next_Permutation {

    public void nextPermutation(int[] nums) {
        int len = nums.length;
        int index;
        int changeIndex;
        int splitIndex = len - 1;
        boolean isSplit = false;
        for(index = len - 1; index > 0; index--) {
            if(nums[index] > nums[index - 1]) {
                isSplit = true;
                changeIndex = index - 1;
                for(int j = len - 1; j > index - 1; j--) {
                    if(nums[j] > nums[changeIndex]) {
                        splitIndex = j;
                        break;
                    }
                }
                int temp;
                temp = nums[splitIndex];
                nums[splitIndex] = nums[changeIndex];
                nums[changeIndex] = temp;
                Arrays.sort(nums, changeIndex + 1, len);
                break;
            }
        }
        if(!isSplit) {
            Arrays.sort(nums);
        }
        for(int n : nums) {
            System.out.print(n + ", ");
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_031_Next_Permutation solution = new No_031_Next_Permutation();
        solution.nextPermutation(new int[] {1, 2, 3});
        solution.nextPermutation(new int[] {3, 2, 1});
        solution.nextPermutation(new int[] {1, 1, 5});
        solution.nextPermutation(new int[] {0, 1, 4, 3, 2});
        solution.nextPermutation(new int[] {0, 2, 4, 3, 1});
    }

}
