package utils;

/**
 * Created by 36191 on 2018/9/22
 */

public class Point {
    public int x;
    public int y;
    public Point() { x = 0; y = 0; }
    public Point(int a, int b) { x = a; y = b; }

    public static Point[] addPoints(int[][] xys) {
        int len = xys.length;
        Point[] ps = new Point[len];
        for (int i = 0; i < len; i++) {
            ps[i] = new Point(xys[i][0], xys[i][1]);
        }
        return ps;
    }
}
