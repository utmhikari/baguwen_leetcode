package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/20
 * 075. 颜色分块排序
 */

public class No_075_Sort_Colors {

    public void swap(int[] nums, int i1, int i2) {
        int temp = nums[i2];
        nums[i2] = nums[i1];
        nums[i1] = temp;
    }

    /*
    void sortColors(int A[], int n) {
        int j = 0, k = n-1;
        for (int i=0; i <= k; i++) {
            if (A[i] == 0)
                swap(A[i], A[j++]);
            else if (A[i] == 2)
                swap(A[i--], A[k--]);
        }
    }
     */
    public void sortColors(int[] nums) {
        int len = nums.length;
        if(len <= 1) {
            return;
        }
        for(int i = 0; i < len - 1; i++) {
            if(nums[i] > nums[i + 1]) {
                swap(nums, i, i + 1);
            }
            if(i != 0 && nums[i] < nums[i - 1]) {
                boolean isChanged = false;
                for(int j = i - 1; j >= 0; j--) {
                    if(nums[j] <= nums[i]) {
                        swap(nums, i, j + 1);
                        isChanged = true;
                        break;
                    }
                }
                if(!isChanged) {
                    swap(nums, i, 0);
                }
                if(nums[i] == 1) {
                    for(int j = i - 1; j >= 0; j--) {
                        if(nums[j] <= nums[i]) {
                            swap(nums, i, j + 1);
                            break;
                        }
                    }
                }
            }
            for(int n : nums) {
                System.out.print(n + ", ");
            }
            System.out.println();
        }
    }



    public static void main(String[] args) {
        No_075_Sort_Colors solution = new No_075_Sort_Colors();
        //solution.sortColors(new int[] {1, 2, 0, 1, 0, 2, 0, 1, 2, 0, 2, 1});
        solution.sortColors(new int[] {1, 2, 2, 2, 2, 0, 0, 0, 1, 1});
    }
}
