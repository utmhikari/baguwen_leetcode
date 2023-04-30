package leetcode_201_300;

import java.util.HashSet;

/**
 * Created by 36191 on 2018/10/2
 */
public class No_217_Contains_Duplicate {
    public boolean containsDuplicate(int[] nums) {
        HashSet<Integer> set = new HashSet<>();
        for (int n : nums) {
            if (set.contains(n)) { return true; }
            set.add(n);
        }
        return false;
    }
}
