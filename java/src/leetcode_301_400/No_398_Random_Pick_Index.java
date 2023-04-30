package leetcode_301_400;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Random;

/**
 * Created by 36191 on 2019/1/26
 */
public class No_398_Random_Pick_Index {
    class Solution {
        
        private int[] primes = new int[] {2, 3, 5, 7, 11};
        
        private int[] numbers;
        private HashMap<String, ArrayList<Integer>> map;
        
        private String encode(int num) {
            StringBuilder sb = new StringBuilder();
            for (int p : primes) { sb.append(num % p); sb.append('_'); }
            return sb.toString();
        }
        
        public Solution(int[] nums) {
            int len = nums.length;
            numbers = new int[len];
            map = new HashMap<>(32);
            for (int i = 0; i < len; i++) {
                numbers[i] = nums[i];
                String code = encode(nums[i]);
                if (!map.containsKey(code)) { map.put(code, new ArrayList<>()); }
                map.get(code).add(i);
            }
        }

        public int pick(int target) {
            String code = encode(target);
            if (!map.containsKey(code)) { return -1; }
            ArrayList<Integer> list = map.get(code);
            int len = list.size();
            Random random = new Random();
            while (true) {
                int rand = random.nextInt(len);
                int idx = list.get(rand);
                if (numbers[idx] == target) { return idx; }
            }
        }
    }
}
