package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/12
 */
public class No_335_Self_Crossing {
    public boolean isSelfCrossing(int[] x) {
        int l = x.length;
        if (l < 4) { return false; }
        for (int i = 3; i < l; i++){
            if( x[i] >= x[i-2] && x[i-1] <= x[i-3]) { return true; }
            if (i >= 4) {
                if (x[i-1] == x[i-3] && x[i] + x[i-4] >= x[i-2]) { return true; }
            }
            if (i >= 5) {
                if (x[i-2] - x[i-4] >= 0 && x[i] >= x[i-2] - x[i-4] && x[i-1] >= x[i-3] - x[i-5] && x[i-1] <= x[i-3]) {
                    return true;
                }
            }
        }
        return false;
    }
    
    public static void main(String[] args) {
        No_335_Self_Crossing solution = new No_335_Self_Crossing();
        int[][] inputs = new int[][] {
                new int[] { 2, 1, 1, 2 },
                new int[] { 1, 2, 3, 4 },
                new int[] { 1, 2, 3, 4, 5, 3, 2 },
        };
        for (int[] input : inputs) {
            System.out.println(solution.isSelfCrossing(input));
        }
    }
    
}
