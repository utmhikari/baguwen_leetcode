package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/8.
 * 026.从排序列表中移除重复项
 */
public class No_026_Remove_Duplicates_Sorted_Array {

    public int removeDuplicates(int[] nums) {
        int count = nums.length > 0 ? 1 : 0;
        for (int n : nums) {
            if (n > nums[count-1]) {
                nums[count++] = n;
            }
        }
        return count;
    }

    public static void main(String[] args) {

    }
}
