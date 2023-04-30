package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/22
 * 080. 去除重复项，最多剩两个
 */

public class No_080_Remove_Duplicates_Sorted_Array_II {

    public int removeDuplicates(int[] nums) {
        int len = nums.length;
        if(len == 0) {
            return 0;
        }
        int count = 1;
        int countDup = 1;
        for(int i = 1; i < len; i++) {
            if(nums[i] == nums[i - 1]) {
                countDup += 1;
                if(countDup <= 2) {
                    nums[count++] = nums[i];
                }
            } else {
                nums[count++] = nums[i];
                countDup = 1;
            }
        }
        return count;
    }

    public static void main(String[] args) {
        No_080_Remove_Duplicates_Sorted_Array_II solution = new No_080_Remove_Duplicates_Sorted_Array_II();
        System.out.println(solution.removeDuplicates(new int[] {1, 1, 1, 1, 2, 2, 2, 3, 3, 4}));
    }

}
