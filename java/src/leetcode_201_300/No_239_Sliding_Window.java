package leetcode_201_300;

import java.util.HashMap;
import java.util.PriorityQueue;

/**
 * Created by 36191 on 2019/8/18
 */
public class No_239_Sliding_Window {
    public int[] maxSlidingWindow(int[] nums, int k) {
        int len = nums.length;
        int[] ans = new int[len + 1 - k];
        if (k == 0 || len == 0) {
            return new int[0];
        }
        ans = new int[len + 1 - k];
        int max = Integer.MIN_VALUE;
        PriorityQueue<Integer> q = new PriorityQueue<>(len, (a, b) -> { return b - a; });
        HashMap<Integer, Integer> cnts = new HashMap<>();
        for (int i = 0; i < k; ++i) {
            q.offer(nums[i]);
            cnts.put(nums[i], cnts.getOrDefault(nums[i], 0) + 1);
        }
        for (int i = k; ; ++i) {
            while (!cnts.containsKey(q.peek())) {
                q.poll();
            }
            ans[i - k] = q.peek();
            int cur = cnts.get(nums[i - k]);
            if (cur == 1) {
                cnts.remove(nums[i - k]);
            } else {
                cnts.put(nums[i - k], cur - 1);
            }
            if (i == len) {
                break;
            } else {
                q.offer(nums[i]);
                cnts.put(nums[i], cnts.getOrDefault(nums[i], 0) + 1);
            }
        }
        return ans;
    }
}
