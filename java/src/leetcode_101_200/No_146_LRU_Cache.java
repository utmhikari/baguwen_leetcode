package leetcode_101_200;

import java.util.HashMap;
import java.util.LinkedList;

/**
 * Created by 36191 on 2019/8/18
 */
public class No_146_LRU_Cache {
    class LRUCache {

        private HashMap<Integer, Integer> cache;
        private LinkedList<Integer> list;
        private int capacity;

        public LRUCache(int capacity) {
            cache = new HashMap<>();
            list = new LinkedList<>();
            this.capacity = capacity;
        }

        public int get(int key) {
            if (!cache.containsKey(key)) {
                return -1;
            }
            list.remove(new Integer(key));
            list.addFirst(key);
            return cache.get(key);
        }

        public void put(int key, int value) {
            if (cache.containsKey(key) || cache.size() < capacity) {
                list.remove(new Integer(key));
            } else {
                cache.remove(list.removeLast());
            }
            list.addFirst(key);
            cache.put(key, value);
        }
    }
}
