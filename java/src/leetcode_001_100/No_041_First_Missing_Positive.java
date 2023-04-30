package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/13
 * 041. 第一个消失的正整数
 */

public class No_041_First_Missing_Positive {

    /*
    // Swap each number to its position!!!!! e.g. 10 to index 9 ~~~
    public int firstMissingPositive(int[] A) {
        int i = 0;
        while(i < A.length){
            if(A[i] == i+1 || A[i] <= 0 || A[i] > A.length) i++;
            else if(A[A[i]-1] != A[i]) swap(A, i, A[i]-1);
            else i++;
        }
        i = 0;
        while(i < A.length && A[i] == i+1) i++;
        return i+1;
    }
    */


    public int firstMissingPositive(int[] nums) {
        int len = nums.length;
        if(len == 0) {
            return 1;
        }
        int misEnd = len + 1;
        int misStart = 1;
        for(int i = 0; i < len; i++) {
            if(nums[i] < misStart) {
                misEnd--;
                nums[i] = Integer.MAX_VALUE;
            }
        }
        while(misEnd > misStart) {
            boolean isChanged = false;
            for(int i = 0; i < len; i++) {
                if(nums[i] == misEnd || nums[i] == misStart || nums[i] < misStart ||
                        (nums[i] > misEnd && nums[i] != Integer.MAX_VALUE)) {
                    if (nums[i] == misStart) {
                        misStart++;
                    } else {
                        misEnd--;
                    }
                    nums[i] = Integer.MAX_VALUE;
                    isChanged = true;
                }
            }
            if(!isChanged) {
                return misStart;
            }
        }
        return misStart;
    }

    public static void main(String[] args) {
        No_041_First_Missing_Positive solution = new No_041_First_Missing_Positive();
        System.out.println(solution.firstMissingPositive(new int[] {2, 1, -1}));
        System.out.println(solution.firstMissingPositive(new int[] {3, 4, -1, 1}));
        System.out.println(solution.firstMissingPositive(new int[] {7, 10, 1, -3, 3, 8, 2, -1, 5, 11, 6, 4}));
        System.out.println(solution.firstMissingPositive(new int[] {2, 1}));
        System.out.println(solution.firstMissingPositive(new int[] {0}));
        System.out.println(solution.firstMissingPositive(new int[] {}));
        System.out.println(solution.firstMissingPositive(new int[] {1, 1}));
        System.out.println(solution.firstMissingPositive(new int[] {2, 2}));
    }
}
