package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/24
 */

public class No_162_Find_Peak_Element {
    private static int _PEAK = Integer.MIN_VALUE;

    private boolean findPeakElement(int[] nums, int idx, int dir, boolean[] hasChecked) {
        System.out.println(idx + ", " + dir);
        if(idx < 0 || idx >= nums.length || hasChecked[idx]) { return false; }
        hasChecked[idx] = true;
        boolean bl = true, br = true;
        if(dir == 0) {
            if (idx != 0) { bl = nums[idx] > nums[idx - 1]; }
            if (idx != nums.length - 1) { br = nums[idx] > nums[idx + 1]; }
        } else if(dir == -1) {
            if (idx == 0) { _PEAK = idx; return true; }
            else { bl = nums[idx] > nums[idx - 1]; }
        } else {
            if (idx == nums.length - 1) { _PEAK = idx; return true;}
            else { br = nums[idx] > nums[idx + 1]; }
        }
        if (bl && br) {
            _PEAK = idx;
            return true;
        } else {
            int l, dl = 0, r, dr = 0;
            if(bl) { l = idx - 2; dl = 0; }
            else { l = idx - 1; dl = -1; }
            if(br) { r = idx + 2; dr = 0; }
            else { r = idx + 1; dr = 1; }
            return findPeakElement(nums, l, dl, hasChecked) || findPeakElement(nums, r, dr, hasChecked);
        }
    }

    public int findPeakElement(int[] nums) {
        _PEAK = Integer.MIN_VALUE;
        boolean[] hasChecked = new boolean[nums.length];
        int m = nums.length / 2;
        findPeakElement(nums, m, 0, hasChecked);
        return _PEAK;
    }
}
