package leetcode_1101_1200;

import java.util.*;

public class No_1192_Critical_Connections {

    private int tarjan(List<List<Integer>> ans, HashMap<Integer, List<Integer>> map,
                       int[] indices, int[] lows, int cur, int prev, int idx) {
        idx++;
        indices[cur] = idx;
        lows[cur] = idx;
        for (int i : map.get(cur)) {
            if (i != prev) {
                if (indices[i] == 0) {
                    idx = tarjan(ans, map, indices, lows, i, cur, idx);
                    lows[cur] = Math.min(lows[i], lows[cur]);
                    if (lows[i] > indices[cur]) {
                        ans.add(new ArrayList<>(Arrays.asList(cur, i)));
                    }
                } else {
                    lows[cur] = Math.min(lows[cur], indices[i]);
                }
            }
        }
        return idx;
    }

    public List<List<Integer>> criticalConnections(int n, List<List<Integer>> connections) {
        List<List<Integer>> ans = new ArrayList<>();
        HashMap<Integer, List<Integer>> map = new HashMap<>(n);
        int[] indices = new int[n];
        int[] lows = new int[n];
        for (List<Integer> c : connections) {
            int u = c.get(0), v = c.get(1);
            if (!map.containsKey(u)) {
                map.put(u, new ArrayList<>());
            }
            if (!map.containsKey(v)) {
                map.put(v, new ArrayList<>());
            }
            map.get(u).add(v);
            map.get(v).add(u);
        }
        int idx = 0;
        for (int i = 0; i < n; i++) {
           if (indices[i] == 0) {
               idx = tarjan(ans, map, indices, lows, i, -1, idx);
           }
        }
        return ans;
    }

    public static void main(String[] args) {
        String s1 = "hello";
        String s2 = "hello";
        System.out.println(s1 == s2);
        System.out.println(s1.equals(s2));
        StringBuilder sb = new StringBuilder();
        for (int i = 1; i < 1000; ++i) {
            sb.append("hello");
        }
        String s3 = sb.toString();
        String s4 = sb.toString();
        System.out.println(s3 == s4);
        System.out.println(s3.equals(s4));
    }
}
