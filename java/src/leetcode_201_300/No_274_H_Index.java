package leetcode_201_300;

import java.util.HashMap;

/**
 * Created by 36191 on 2018/10/21
 */
public class No_274_H_Index {
    public int hIndex(int[] citations) {
        int l = citations.length, h = l, minus = 0;
        HashMap<Integer, Integer> searched = new HashMap<>(32);
        for (int i = 0; i < l; i++) {
            if (citations[i] < h) {
                ++minus;
                if (searched.containsKey(citations[i])) { searched.put(citations[i], searched.get(citations[i]) + 1); }
                else { searched.put(citations[i], 1); }
                if (l - minus - h < 0) { --h; }
                if (searched.containsKey(h)) { minus -= searched.get(h); searched.remove(h); }
            }
        }
        return h;
    }
}
