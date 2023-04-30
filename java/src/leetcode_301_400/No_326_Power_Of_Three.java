package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/12
 */
public class No_326_Power_Of_Three {
    public boolean isPowerOfThree(int n) {
        if (n == 0) { return false; }
        while (n % 3 == 0) { n /= 3; }
        return n == 1;
    }
}
