package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/10.
 * 033.在转序后的数组中搜索
 */
public class No_033_Search_In_Rotated_Sorted_Array {

    public int findMinIndex(int[] nums, int startIndex, int endIndex) {
        if(startIndex == endIndex) {
            return -1;
        }
        int half = startIndex + (endIndex - startIndex) / 2;
        if(nums[half] > nums[half + 1]) {
            return half + 1;
        } else {
            int former = findMinIndex(nums, startIndex, half);
            if(former != -1) {
                return former;
            }
            int latter = findMinIndex(nums, half + 1, endIndex);
            if(latter != -1) {
                return latter;
            }
        }
        return -1;
    }

    public int binarySearch(int[] nums, int target, int startIndex, int endIndex) {
        if(startIndex == endIndex) {
            return nums[startIndex] == target ? startIndex : -1;
        }
        int half = startIndex + (endIndex - startIndex) / 2;
        if(target > nums[half]) {
            return binarySearch(nums, target, half + 1, endIndex);
        } else {
            return binarySearch(nums, target, startIndex, half);
        }
    }

    public int search(int[] nums, int target) {
        int len = nums.length;
        if(len == 0) {
            return -1;
        }
        if(len == 1) {
            return nums[0] == target ? 0 : -1;
        }
        int minIndex = findMinIndex(nums, 0, len - 1);
        if(minIndex == -1) {
            return binarySearch(nums, target, 0, len - 1);
        } else {
            if(target > nums[len - 1]) {
                return binarySearch(nums, target, 0, minIndex - 1);
            } else {
                return binarySearch(nums, target, minIndex, len - 1);
            }
        }
    }

    public static void main(String[] args) {
        No_033_Search_In_Rotated_Sorted_Array solution = new No_033_Search_In_Rotated_Sorted_Array();
        System.out.println(solution.search(new int[] {4, 5, 6, 7, 0, 1, 2}, 6));
        System.out.println(solution.search(new int[] {4, 5, 6, 7, 0, 1, 2}, 3));
        System.out.println(solution.search(new int[] {4, 5, 6, 7, 0, 1, 2}, 0));
    }

}
