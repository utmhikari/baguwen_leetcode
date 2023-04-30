package leetcode_301_400;

import java.util.HashMap;

/**
 * Created by 36191 on 2019/2/4
 */
public class No_397_Integer_Replacement {
    private HashMap<Long, Integer> visit;

    private void rep(long n, int d) {
        if (!visit.containsKey(n) || visit.get(n) > d) {
            visit.put(n, d);
            if (n != 1) {
                if (n % 2 == 0) {
                    rep(n / 2, d + 1);
                } else {
                    rep(n + 1, d + 1);
                    rep(n - 1, d + 1);
                }
            }
        }
    }

    public int integerReplacement(int n) {
        visit = new HashMap<>();
        rep((long)n, 0);
        return visit.get(1L);
    }
}
