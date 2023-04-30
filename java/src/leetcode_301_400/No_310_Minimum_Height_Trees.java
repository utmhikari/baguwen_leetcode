package leetcode_301_400;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by 36191 on 2018/11/27
 */
public class No_310_Minimum_Height_Trees {
    public List<Integer> findMinHeightTrees(int n, int[][] edges) {
        List<Integer> ans = new ArrayList<>();
        if (n <= 1) { ans.add(0); return ans; }
        if (n == 2) { ans.add(0); ans.add(1); return ans; }
        int[] degrees = new int[n];
        ArrayList<ArrayList<Integer>> graph = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            graph.add(new ArrayList<>());
        }
        for (int[] edge : edges) {
            degrees[edge[0]]++;
            degrees[edge[1]]++;
            graph.get(edge[1]).add(edge[0]);
            graph.get(edge[0]).add(edge[1]);
        }
        ArrayList<Integer> queue = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            if (degrees[i] == 1) { queue.add(i); }
        }
        while (!queue.isEmpty()) {
            ArrayList<Integer> temp = new ArrayList<>();
            if (!ans.isEmpty()) { ans.clear(); }
            ans.addAll(queue);
            for (int i : queue) {
                for (int j : graph.get(i)) {
                    degrees[j]--;
                    if (degrees[j] == 1) { temp.add(j); }
                }
            }
            if (!queue.isEmpty()) { queue.clear(); }
            queue.addAll(temp);
        }
        return ans;
    }

    public static void main(String[] args) {
        No_310_Minimum_Height_Trees solution = new No_310_Minimum_Height_Trees();
        List<Integer> ans1 = solution.findMinHeightTrees(6, new int[][] {
                new int[] {0, 3},
                new int[] {1, 3},
                new int[] {2, 3},
                new int[] {3, 4},
                new int[] {5, 4},
        });
        for (int i : ans1) {
            System.out.print(i + ", ");
        }
        System.out.println();
        List<Integer> ans2 = solution.findMinHeightTrees(4, new int[][] {
                new int[] {1, 0},
                new int[] {1, 2},
                new int[] {3, 1},
        });
        for (int i : ans2) {
            System.out.print(i + ", ");
        }
        System.out.println();
    }
}
