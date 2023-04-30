package leetcode_401_500;

/**
 * Created by 36191 on 2019/2/24
 */
public class No_441_Arranging_Coins {
    public int arrangeCoins(int n) {
        int s = 0;
        while (n >= 0) { n -= ++s; }
        return s - 1;
    }
}
