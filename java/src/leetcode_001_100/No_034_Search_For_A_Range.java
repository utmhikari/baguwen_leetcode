package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/10.
 * 034.搜索一个相等数的范围
 */
public class No_034_Search_For_A_Range {

    public int binarySearchDiff(int[] nums, int target, int startIndex, int endIndex, int direction) {
        if(startIndex > endIndex) {
            return endIndex;
        }
        if(startIndex == endIndex) {
            if(direction == 1) {
                return nums[endIndex] > target ? (endIndex - 1) : endIndex;
            } else if(direction == -1) {
                return nums[endIndex] < target ? (endIndex + 1) : endIndex;
            } else {
                return -1;
            }
        }
        int half = startIndex + (endIndex - startIndex) / 2;
        if(direction == 1) {
            if(nums[half] > target) {
                return binarySearchDiff(nums, target, startIndex, half - 1, 1);
            } else {
                return binarySearchDiff(nums, target, half + 1, endIndex, 1);
            }
        } else if(direction == -1) {
            if(nums[half] < target) {
                return binarySearchDiff(nums, target, half + 1, endIndex, -1);
            } else {
                return binarySearchDiff(nums, target, 0, half, -1);
            }
        } else {
            return -1;
        }
    }

    public int[] binarySearch(int[] nums, int target, int startIndex, int endIndex) {
        if(startIndex > endIndex) {
            return new int[] {-1, -1};
        }
        if(startIndex == endIndex) {
            return nums[startIndex] == target ? new int[] {startIndex, startIndex} : new int[] {-1, -1};
        }
        int half = startIndex + (endIndex - startIndex) / 2;
        if(target > nums[half]) {
            return binarySearch(nums, target, half + 1, endIndex);
        } else if(target < nums[half]) {
            return binarySearch(nums, target, 0, half - 1);
        } else {
            int[] ans = new int[] {-1, -1};
            ans[0] = binarySearchDiff(nums, target, startIndex, half, -1);
            ans[1] = binarySearchDiff(nums, target, half, endIndex, 1);
            return ans;
        }
    }

    public int[] searchRange(int[] nums, int target) {
        int len = nums.length;
        if(len == 0) {
            return new int[] {-1, -1};
        }
        if(len == 1) {
            return nums[0] == target ? new int[] {0, 0} : new int[] {-1, -1};
        }
        return binarySearch(nums, target, 0, len - 1);
    }

    public void output(int[] nums, int target) {
        int[] ans = searchRange(nums, target);
        for(int i : ans) {
            System.out.print(i + ", ");
        }
        System.out.println();
    }


    public static void main(String[] args) {
        No_034_Search_For_A_Range solution = new No_034_Search_For_A_Range();
        solution.output(new int[] {5, 7, 7, 8, 8, 10}, 8);
        solution.output(new int[] {5, 7, 7, 8, 8, 10}, 6);
        solution.output(new int[] {2, 2}, 1);
        solution.output(new int[] {5, 5, 5, 5, 5, 5}, 5);
    }


}
