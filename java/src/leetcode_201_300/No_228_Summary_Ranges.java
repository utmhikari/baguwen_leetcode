package leetcode_201_300;

import java.util.List;

/**
 * Created by 36191 on 2018/10/15
 */

import java.util.*;

public class No_228_Summary_Ranges {
    public List<String> summaryRanges(int[] nums) {
        List<String> ans = new ArrayList<>();
        int l = nums.length;
        if (l == 0) { return ans; }
        List<int[]> pairs = new ArrayList<>();
        int start = 0, i = 1;
        while (i < l) {
            while (nums[i] - nums[i++ - 1] == 1) { if (i == l) { break; } }
            pairs.add(new int[] {nums[start], nums[i - 1]});
            start = i++;
        }
        if (start == l - 1) { pairs.add(new int[] {nums[l - 1], nums[l - 1]}); }
        for (int[] pair : pairs) {
            if (pair[0] == pair[1]) { ans.add(String.valueOf(pair[0])); }
            else { ans.add(String.valueOf(pair[0]) + "->" + String.valueOf(pair[1])); }
        }
        return ans;
    }
}
