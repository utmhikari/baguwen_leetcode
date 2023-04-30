package leetcode_201_300;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.LinkedList;

/**
 * Created by 36191 on 2018/10/5
 */


public class No_210_Course_Schedule_II {

    private static LinkedList<Integer> orderList = new LinkedList<>();

    private boolean findCircle(int i, ArrayList<ArrayList<Integer>> nexts, HashSet<Integer> removes,
                               HashSet<Integer> steps) {
        if (removes.contains(i)) { return false; }
        if (nexts.get(i).size() == 0) { removes.add(i); orderList.addFirst(i); return false; }
        if (steps.contains(i)) { return true; }
        steps.add(i);
        for (int next : nexts.get(i)) {
            if (findCircle(next, nexts, removes, steps)) { return true; }
        }
        orderList.addFirst(i);
        steps.remove(i);
        removes.add(i);
        return false;
    }

    public int[] findOrder(int numCourses, int[][] prerequisites) {
        ArrayList<ArrayList<Integer>> nexts = new ArrayList<>();
        for (int i = 0; i < numCourses; i++) { nexts.add(new ArrayList<>()); }
        for (int[] p : prerequisites) { nexts.get(p[1]).add(p[0]); }
        HashSet<Integer> removes = new HashSet<>();
        if (!orderList.isEmpty()) { orderList.clear(); }
        for (int i = 0; i < numCourses; i++) {
            if (!removes.contains(i) && findCircle(i, nexts, removes, new HashSet<>())) { return new int[0]; }
        }
        int[] order = new int[numCourses];
        for (int i = 0; i < numCourses; i++) { order[i] = orderList.removeFirst(); }
        return order;
    }

    public static void main(String[] args) {
        No_210_Course_Schedule_II solution = new No_210_Course_Schedule_II();
        ArrayList<int[][]> inputs = new ArrayList<int[][]>() {{
            add(new int[][] {
                    new int[] {2, 3},
                    new int[] {3, 0},
                    new int[] {1, 4},
                    new int[] {4, 0},
                    //new int[] {0, 1}
            });
        }};

        for (int i : solution.findOrder(5, inputs.get(0))) {
            System.out.print(i + ", ");
        }
        System.out.println();

    }
}
