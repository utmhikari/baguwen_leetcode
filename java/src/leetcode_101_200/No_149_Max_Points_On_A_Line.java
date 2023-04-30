package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/22
 * No.149 最多点的一条线
 */
import utils.*;
import java.util.*;

public class No_149_Max_Points_On_A_Line {

    static int max;

    private int gcd(int a, int b) {
        return b == 0 ? a : gcd(b, a % b);
    }

    private void addVal(HashMap<String, HashSet<Integer>> kpi, String k, int j) {
        if (!kpi.containsKey(k)) {
            kpi.put(k, new HashSet<>());
            kpi.get(k).add(j);
            max = Math.max(max, 1);
        } else if (!kpi.get(k).contains(j)) {
            kpi.get(k).add(j);
            if(kpi.get(k).size() > max) {
                System.out.print(k + ": ");
                for (int i : kpi.get(k)) { System.out.print(i + ", "); }
                System.out.println();
            }
            max = Math.max(max, kpi.get(k).size());

        }
    }

    public int maxPoints(Point[] points) {
        if (points == null || points.length == 0) { return 0; }
        if (points.length == 1) { return 1; }
        // <k, point indices>
        HashMap<String, HashSet<Integer>> kpi = new HashMap<>(32);
        max = 0;
        for(int i = 0; i < points.length - 1; i++) {
            for (int j = i + 1; j < points.length; j++) {
                int x = points[j].x - points[i].x, y = points[j].y - points[i].y;
                String inf = "INF_" + String.valueOf(points[j].x), zero = "0_" + String.valueOf(points[j].y);
                if (x == 0 && y == 0) {
                    for (String k : kpi.keySet()) {
                        if (!k.startsWith("I") && !k.startsWith("0")) {
                            String[] yxb = k.split("_");
                            int y0 = Integer.parseInt(yxb[0]), x0 = Integer.parseInt(yxb[1]), b = Integer.parseInt(yxb[2]);
                            if (x0 * points[i].y - y0 * points[i].x / x0 == b) {
                                addVal(kpi, k, i); addVal(kpi, k, j);
                            }
                        }
                    }
                    addVal(kpi, inf , i); addVal(kpi, inf, j);
                    addVal(kpi, zero , i); addVal(kpi, zero, j);
                }
                else if (x == 0) { addVal(kpi, inf, i); addVal(kpi, inf, j); }
                else if (y == 0) { addVal(kpi, zero, i); addVal(kpi, zero, j); }
                else {
                    int g = gcd(x, y), y0 = y / g, x0 = x / g, b = (x0 * points[i].y - y0 * points[i].x) / x0;
                    String k = String.valueOf(y0) + "_" + String.valueOf(x0) + "_" + String.valueOf(b);
                    addVal(kpi, k, i); addVal(kpi, k, j);
                }
            }
        }
        return max;
    }

    public static void main(String[] args) {
        No_149_Max_Points_On_A_Line solution = new No_149_Max_Points_On_A_Line();
        int[][] xys1 = new int[][] {
                new int[] {1, 1}, new int[] {3, 2}, new int[] {5, 3},
                new int[] {4, 1}, new int[] {2, 3}, new int[] {1, 4}
        };
        Point[] ps1 = Point.addPoints(xys1);
        System.out.println(solution.maxPoints(ps1));
        int[][] xys2 = new int[][] {
                new int[] {1, 1}, new int[] {1, 1}, new int[] {2, 3}
        };
        Point[] ps2 = Point.addPoints(xys2);
        System.out.println(solution.maxPoints(ps2));
        int[][] xys3 = new int[][] {
                new int[] {0, 12}, new int[] {5, 2}, new int[] {2, 5},
                new int[] {0, -5}, new int[] {1, 5}, new int[] {2, -2},
                new int[] {5, -4}, new int[] {3, 4}, new int[] {-2, 4},
                new int[] {-1, 4}, new int[] {0, -5}, new int[] {0, -8},
                new int[] {-2, 1}, new int[] {0, -11}, new int[] {0, -9}
        };
        Point[] ps3 = Point.addPoints(xys3);
        System.out.println(solution.maxPoints(ps3));
    }
}
