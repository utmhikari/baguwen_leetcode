package leetcode_301_400;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Random;

/**
 * Created by 36191 on 2019/2/9
 */
public class No_380_Insert_Delete_GetRandom_O1 {

    class RandomizedSet {
        private HashMap<Integer, Integer> map;
        private List<Integer> list;
        private int left, right;

        /** Initialize your data structure here. */
        public RandomizedSet() {
            map = new HashMap<>(16);
            list = new ArrayList<>();
            left = 0;
            right = 0;
        }

        /** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
        public boolean insert(int val) {
            if (map.containsKey(val)) { return false; }
            list.add(val);
            map.put(val, right++);
            return true;
        }

        /** Removes a value from the set. Returns true if the set contained the specified element. */
        public boolean remove(int val) {
            if (!map.containsKey(val)) { return false; }
            int idx = map.remove(val);
            if (idx != left) {
                int head = list.get(left);
                list.set(idx, head);
                map.put(head, idx);
            }
            ++left;
            return true;
        }

        /** Get a random element from the set. */
        public int getRandom() {
            Random random = new Random();
            int nextInt = random.nextInt(right - left);
            return list.get(nextInt + left);
        }
    }
}
