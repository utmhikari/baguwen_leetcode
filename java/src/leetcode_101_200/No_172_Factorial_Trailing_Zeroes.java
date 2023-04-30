package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/25
 */

public class No_172_Factorial_Trailing_Zeroes {
    public int trailingZeroes(int n) {
        int sum = 0;
        while (n >= 5) { n = n / 5; sum += n;}
        return sum;
    }
}
