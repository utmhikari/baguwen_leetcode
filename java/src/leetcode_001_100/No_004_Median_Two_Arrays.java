package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/3.
 * 004.两个排序过的数组的中位数
 */
public class No_004_Median_Two_Arrays {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        double median = 0.0;
        int length1 = nums1.length;
        int length2 = nums2.length;
        int length = length1 + length2;
        int size = length / 2 + 1;
        int[] nums = new int[size];
        int i = 0;
        int i1 = 0;
        int i2 = 0;
        while(i < size) {
            if(i1 < length1 && i2 < length2) {
                if(nums1[i1] < nums2[i2]) {
                    nums[i] = nums1[i1];
                    i1++;
                } else {
                    nums[i] = nums2[i2];
                    i2++;
                }
            } else if(i1 < length1 && i2 >= length2) {
                nums[i] = nums1[i1];
                i1++;
            } else if(i2 < length2 && i1 >= length1) {
                nums[i] = nums2[i2];
                i2++;
            }
            i++;
        }
        if(length % 2 == 0) {
            median = (double)(nums[size - 1] + nums[size - 2]) / 2;
        } else {
            median = (double)nums[size - 1];
        }
        return median;
    }

    public void output(double median) {
        String median_str = String.format("%.1f", median);
        System.out.println("The Median is: " + median_str);
    }

    public static void main(String[] args) {
        No_004_Median_Two_Arrays solution = new No_004_Median_Two_Arrays();
        int[] nums1 = {1, 3};
        int[] nums2 = {2};
        solution.output(solution.findMedianSortedArrays(nums1, nums2));
        int[] nums3 = {1, 2};
        int[] nums4 = {3, 4};
        solution.output(solution.findMedianSortedArrays(nums3, nums4));
    }
}
