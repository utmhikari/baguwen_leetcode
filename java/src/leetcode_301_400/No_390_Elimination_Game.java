package leetcode_301_400;

/**
 * Created by 36191 on 2019/2/4
 */
public class No_390_Elimination_Game {
    public int lastRemaining(int n) {
        boolean left = true;
        int rem = n, cur = 1, step = 1;
        while (rem > 1) {
            if (left || rem % 2 == 1) { cur += step; }
            rem /= 2; step *= 2;
            left = !left;
        }
        return cur;
    }
}
