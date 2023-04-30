package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/21
 * No.135 Candy
 * 每个人至少发一个糖果，邻居里边大的要发多，最少总共几个呢
 */
import java.util.*;

public class No_135_Candy {

    private int f(int n) {
        return n * (n + 1) / 2;
    }

    public int candy(int[] ratings) {
        if (ratings.length == 1) { return 1; }
        if (ratings.length == 2) { return ratings[0] == ratings[1] ? 2 : 3; }
        long sum = 0;
        int i = 0, len = ratings.length, n, t = 1;
        while(i < len) {
            n = 1;
            while (i + 1 < len && ratings[i + 1] < ratings[i]) {
                i++;
                n++;
            }
            if (n == 1) {
                if (i == 0 || ratings[i] == ratings[i - 1]) { t = 1; }
                else { t++; }
                System.out.print(t + ", ");
                sum += t;
            } else {
                int add = f(n);
                if (t >= n && ratings[i - n] < ratings[i - n + 1]) { add += t + 1 - n; }
                t = 1;
                System.out.print(add + ", ");
                sum += add;
            }
            i++;
        }
        System.out.println();
        return (int)sum;
    }

    public static void main(String[] args) {
        No_135_Candy solution = new No_135_Candy();
        int[][] rs = new int[][] {
                new int[] {1, 0, 2},
                new int[] {1, 2, 2},
                new int[] {1, 4, 3, 2, 2, 1, 4, 5, 6, 1},
                new int[] {1, 2, 4, 4, 4, 3},
                new int[] {0, 1, 2, 3, 2, 1},
                new int[] {1, 2, 3, 3, 2, 1},
                new int[] {1, 2, 4, 4, 3}
        };
        for (int[] r : rs) { System.out.println(solution.candy(r)); }
    }
}
