package leetcode_201_300;

import java.util.List;

/**
 * Created by 36191 on 2018/10/15
 */

import java.util.*;

public class No_229_Majority_Element_II {
    public List<Integer> majorityElement(int[] nums) {
        List<Integer> ans = new ArrayList<>();
        int l = nums.length;
        if (l == 0) { return ans; }
        int[] marks = new int[] {nums[0], nums[0], 0, 0};
        for (int i = 0; i < l; i++) {
            if (nums[i] == marks[0]) { marks[2]++; }
            else if (nums[i] == marks[1]) { marks[3]++; }
            else if (marks[2] == 0) { marks[0] = nums[i]; marks[2] = 1; }
            else if (marks[3] == 0) { marks[1] = nums[i]; marks[3] = 1; }
            else { marks[2]--; marks[3]--; }
        }
        marks[2] = 0; marks[3] = 0;
        for (int i = 0; i < l; i++) {
            if (nums[i] == marks[0]) { marks[2]++; }
            else if (nums[i] == marks[1]) { marks[3]++; }
        }
        if (marks[2] > l / 3) { ans.add(marks[0]); }
        if (marks[3] > l / 3) { ans.add(marks[1]); }
        return ans;
    }
}
