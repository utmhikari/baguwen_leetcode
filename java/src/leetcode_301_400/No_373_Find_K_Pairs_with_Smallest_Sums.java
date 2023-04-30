package leetcode_301_400;

import java.util.*;

/**
 * Created by 36191 on 2019/2/4
 */
public class No_373_Find_K_Pairs_with_Smallest_Sums {

    public List<int[]> kSmallestPairs(int[] nums1, int[] nums2, int k) {
        List<int[]> ans = new ArrayList<>();
        int l1 = nums1.length, l2 = nums2.length;
        Arrays.sort(nums1); Arrays.sort(nums2);
        if (k == 0 || l1 == 0 || l2 == 0) { return ans; }
        PriorityQueue<int[]> pq = new PriorityQueue<>((p1, p2) -> p1[0] + p1[1] - p2[0] - p2[1]);
        for (int i = 0; i < l1 && i < k; i++) {
            pq.offer(new int[] { nums1[i], nums2[0], 0 });
        }
        while (k-- > 0 && !pq.isEmpty()) {
            int[] cur = pq.poll();
            ans.add(new int[] { cur[0], cur[1] });
            if (cur[2] != l2 - 1) { pq.offer(new int[] { cur[0], nums2[cur[2] + 1], cur[2] + 1 }); }
        }
        return ans;
    }
}
