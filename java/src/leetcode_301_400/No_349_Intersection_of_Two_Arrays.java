package leetcode_301_400;

import java.util.HashSet;
import java.util.Iterator;

/**
 * Created by 36191 on 2019/1/13
 */
public class No_349_Intersection_of_Two_Arrays {
    public int[] intersection(int[] nums1, int[] nums2) {
        HashSet<Integer> s1 = new HashSet<Integer>(), ans = new HashSet<Integer>();
        for (int i : nums1) { s1.add(i); }
        for (int i : nums2)  { if (s1.contains(i)) { ans.add(i); }}
        Iterator<Integer> iter = ans.iterator();
        int[] a = new int[ans.size()];
        for (int i = 0; i < a.length; i++) { a[i] = iter.next(); }
        return a;
    }
}
