package leetcode_301_400;

import java.util.Random;

/**
 * Created by 36191 on 2019/2/9
 */
public class No_384_Shuffle_An_Array {
    class Solution {

        private int[] ori;

        public Solution(int[] nums) {
            int len = nums.length;
            ori = new int[len];
            System.arraycopy(nums, 0, ori, 0, len);
        }

        /** Resets the array to its original configuration and return it. */
        public int[] reset() {
            return ori;
        }

        /** Returns a random shuffling of the array. */
        public int[] shuffle() {
            Random random = new Random();
            int len = ori.length;
            int[] sf = new int[len], order = new int[len];
            for (int i = 0; i < len; i++) { order[i] = i; }
            for (int i = 0; i < len; i++) {
                int randIdx = i + random.nextInt(len - i);
                sf[i] = ori[order[randIdx]];
                if (order[randIdx] != order[i]) { order[randIdx] = order[i]; }
            }
            return sf;
        }
    }
}
