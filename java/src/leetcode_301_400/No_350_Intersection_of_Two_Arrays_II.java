package leetcode_301_400;

import java.util.ArrayList;
import java.util.HashMap;

/**
 * Created by 36191 on 2019/1/13
 */
public class No_350_Intersection_of_Two_Arrays_II {
    public int[] intersect(int[] nums1, int[] nums2) {
        if (nums1.length == 0 || nums2.length == 0) { return new int[0]; }
        HashMap<Integer, Integer> map1 = new HashMap<>(), map2 = new HashMap<>();
        for (int i : nums1) { if (map1.containsKey(i)) { map1.put(i, map1.get(i) + 1); } else { map1.put(i, 1); } }
        for (int i : nums2) { if (map2.containsKey(i)) { map2.put(i, map2.get(i) + 1); } else { map2.put(i, 1); } }
        ArrayList<Integer> ans = new ArrayList<>();
        for (int i : map1.keySet()) {
            if (map2.containsKey(i)) {
                int times = Math.min(map1.get(i), map2.get(i));
                while (times-- != 0) { ans.add(i); }
            }
        }
        if (ans.size() == 0) { return new int[0]; }
        int[] a = new int[ans.size()];
        for (int i = 0; i < a.length; i++) { a[i] = ans.get(i); }
        return a;
    }
}
