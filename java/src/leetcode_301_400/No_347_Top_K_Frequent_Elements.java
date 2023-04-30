package leetcode_301_400;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Created by 36191 on 2019/1/13
 */
public class No_347_Top_K_Frequent_Elements {
    public List<Integer> topKFrequent(int[] nums, int k) {
        List<Integer> ans = new ArrayList<>();
        HashMap<Integer, Integer> count = new HashMap<>();
        for (int i : nums) {
            count.put(i, count.getOrDefault(i, 0) + 1);
        }
        List<Map.Entry<Integer, Integer>> list = new ArrayList<>(count.entrySet());
        list.sort(Map.Entry.comparingByValue());
        for(int i = list.size() - 1; i >= list.size() - k; i--) {
            ans.add(list.get(i).getKey());
        }
        return ans;
    }
}
