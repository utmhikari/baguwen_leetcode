package leetcode_401_500;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Created by 36191 on 2019/2/24
 */
public class No_447_Number_of_Boomerangs {
    public int numberOfBoomerangs(int[][] points) {
        HashMap<Integer, List<Integer>> map;
        int l = points.length, ans = 0;
        for (int i = 0; i < l; ++i) {
            map = new HashMap<>();
            for (int j = 0; j < l; ++j) {
                if (j != i) {
                    int dx = points[i][0] - points[j][0], dy = points[i][1] - points[j][1];
                    int dz = dx * dx + dy * dy;
                    if (!map.containsKey(dz)) { map.put(dz, new ArrayList<>()); }
                    map.get(dz).add(j);
                }
            }
            for (Map.Entry<Integer, List<Integer>> e : map.entrySet()) {
                List<Integer> list = e.getValue();
                int size = list.size();
                ans += size * (size - 1);
            }
        }
        return ans;
    }
}
