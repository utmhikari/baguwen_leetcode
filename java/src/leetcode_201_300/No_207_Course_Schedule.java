package leetcode_201_300;

import java.util.ArrayList;
import java.util.HashSet;

/**
 * Created by 36191 on 2018/10/5
 */
public class No_207_Course_Schedule {

    private boolean findCircle(int i, ArrayList<ArrayList<Integer>> nexts, HashSet<Integer> removes,
                               HashSet<Integer> steps) {
        if (removes.contains(i)) { return false; }
        if (nexts.get(i).size() == 0) { removes.add(i); return false; }
        if (steps.contains(i)) { return true; }
        steps.add(i);
        for (int next : nexts.get(i)) {
            if (findCircle(next, nexts, removes, steps)) { return true; }
        }
        steps.remove(i);
        removes.add(i);
        return false;
    }

    public boolean canFinish(int numCourses, int[][] prerequisites) {
        ArrayList<ArrayList<Integer>> nexts = new ArrayList<>();
        for (int i = 0; i < numCourses; i++) { nexts.add(new ArrayList<>()); }
        for (int[] p : prerequisites) { nexts.get(p[1]).add(p[0]); }
        HashSet<Integer> removes = new HashSet<>();
        for (int i = 0; i < numCourses; i++) {
            if (!removes.contains(i) && findCircle(i, nexts, removes, new HashSet<>())) { return false; }
        }
        return true;
    }

    public static void main(String[] args) {
        No_207_Course_Schedule solution = new No_207_Course_Schedule();
        ArrayList<int[][]> inputs = new ArrayList<int[][]>() {{
            add(new int[][] {
                new int[] {2, 3},
                new int[] {3, 0},
                new int[] {1, 4},
                new int[] {4, 0},
                //new int[] {0, 1}
            });
        }};

        System.out.println(solution.canFinish(5, inputs.get(0)));
    }
}
