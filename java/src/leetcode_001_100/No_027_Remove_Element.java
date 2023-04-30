package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/9.
 * 027.移除元素
 */
public class No_027_Remove_Element {

    public int removeElement(int[] nums, int val) {
        int count = 0;
        int len = nums.length;
        for (int i = 0; i < len; i++) {
            if(nums[i] != val) {
                nums[count++] = nums[i];
            }
        }
        return count;
    }

    public static void main(String[] args) {
        No_027_Remove_Element solution = new No_027_Remove_Element();
        solution.removeElement(new int[]{3, 2, 2, 3}, 3);
        solution.removeElement(new int[]{1}, 1);
    }
}
