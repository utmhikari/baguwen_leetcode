package leetcode_201_300;

import java.util.ArrayList;
import java.util.HashMap;

/**
 * Created by 36191 on 2018/10/2
 */
public class No_219_Contains_Duplicate_II {

    public boolean containsNearbyDuplicate(int[] nums, int k) {
        if (nums == null || nums.length == 0 || k < 1) { return false; }
        int len = nums.length;
        HashMap<Integer, ArrayList<Integer>> map = new HashMap<>(len);
        for (int i = 0; i < len; i++) {
            if (map.containsKey(nums[i])) {
                ArrayList<Integer> tmp = map.get(nums[i]);
                if (i - tmp.get(tmp.size() - 1) <= k) { return true; }
                else { tmp.add(i); }
            } else {
                map.put(nums[i], new ArrayList<>());
                map.get(nums[i]).add(i);
            }
        }
        return false;
    }
}
