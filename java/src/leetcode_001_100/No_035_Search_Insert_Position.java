package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/10.
 * 035.搜索插入位置
 */

public class No_035_Search_Insert_Position {

    public int binarySearchInsert(int[] nums, int target, int startIndex, int endIndex) {
        if(startIndex == endIndex) {
            if(nums[startIndex] < target) {
                return startIndex + 1;
            } else {
                return startIndex;
            }
        } else {
            int half = startIndex + (endIndex - startIndex) / 2;
            if(nums[half] == target) {
                return half;
            } else if(nums[half] < target) {
                return binarySearchInsert(nums, target, half + 1, endIndex);
            } else {
                return binarySearchInsert(nums, target, startIndex, half);
            }
        }
    }

    public int searchInsert(int[] nums, int target) {
        int len = nums.length;
        if(len == 0) {
            return 0;
        }
        return binarySearchInsert(nums, target, 0, len - 1);
    }

    public void output(int[] nums, int target) {
        System.out.println(searchInsert(nums, target));
    }


    public static void main(String[] args) {
        No_035_Search_Insert_Position solution = new No_035_Search_Insert_Position();
        int[] testNums = new int[] {1, 3, 5, 6};
        solution.output(testNums, 5);
        solution.output(testNums, 2);
        solution.output(testNums, 7);
        solution.output(testNums, 0);
    }
}