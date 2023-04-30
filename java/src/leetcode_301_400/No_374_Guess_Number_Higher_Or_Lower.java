package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/16
 */
public class No_374_Guess_Number_Higher_Or_Lower {
    private int guess(int num) { return 1; }
    
    private int g(int l, int r) {
        if (l >= r) { return l; }
        int m = l + (r - l) / 2, a = guess(m);
        if (a == 0) { return m; }
        else if (a == 1) { return g(m + 1, r); }
        else { return g(l, m - 1); }
    }
    public int guessNumber(int n) {
        return g(1, n);
    }
}
