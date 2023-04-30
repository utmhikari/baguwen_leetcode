package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/14
 */
public class No_223_Rectangle_Area {

    public int computeArea(int A, int B, int C, int D, int E, int F, int G, int H) {
        int sum = (C - A) * (D - B) + (G - E) * (H - F);
        int x1 = Math.max(A, E), y1 = Math.max(B, F), x2 = Math.min(C, G), y2 = Math.min(D, H);
        if (x2 <= x1 || y2 <= y1) { return sum; }
        else { return sum - (x2 - x1) * (y2 - y1); }
    }

}
