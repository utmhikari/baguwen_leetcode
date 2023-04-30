package leetcode_201_300;

import java.util.List;

/**
 * Created by 36191 on 2018/10/7
 */
import java.util.*;

public class No_216_Combination_Sum_III {

    private void find(int k, int n, List<Integer> ans, List<List<Integer>> list) {
        if (!(n > 9 * k || n < k)) {
            int start = ans.isEmpty() ? 1 : ans.get(ans.size() - 1) + 1;
            if (k == 1) {
                if (n >= start) {
                    ans.add(n);
                    list.add(new ArrayList<>(ans));
                    ans.remove(ans.size() - 1);
                }
            } else {
                for (int i = start; i <= 9; i++) {
                    ans.add(i);
                    find(k - 1, n - i, ans, list);
                    ans.remove(ans.size() - 1);
                }
            }
        }
    }

    public List<List<Integer>> combinationSum3(int k, int n) {
        List<List<Integer>> list = new ArrayList<>();
        List<Integer> ans = new ArrayList<>();
        find(k, n, ans, list);
        return list;
    }
}
