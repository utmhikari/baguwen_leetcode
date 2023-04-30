package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/29
 * 088. 合并排序后列表
 */

public class No_088_Merge_Sorted_Array {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int _m = 0;
        int _n = 0;
        while(_m < m && _n < n) {
            if(nums2[_n] < nums1[_m + _n]) {
                for(int i = m + _n - 1; i >= _m + _n; i--) {
                    nums1[i + 1] = nums1[i];
                }
                nums1[_m + _n] = nums2[_n];
                _n++;
            } else {
                _m++;
            }
        }
        if(_n < n) {
            for(int i = m + _n; i < m + n; i++) {
                nums1[i] = nums2[i - m];
            }
        }
        for(int num : nums1) {
            System.out.println(num);
        }
    }

    public static void main(String[] args) {
        No_088_Merge_Sorted_Array solution = new No_088_Merge_Sorted_Array();
        int[] nums1 = new int[9];
        nums1[0] = 1;
        nums1[1] = 3;
        nums1[2] = 6;
        int[] nums2 = new int[4];
        nums2[0] = 2;
        nums2[1] = 4;
        nums2[2] = 5;
        nums2[3] = 7;
        solution.merge(nums1, 3, nums2, 4);
    }
}
