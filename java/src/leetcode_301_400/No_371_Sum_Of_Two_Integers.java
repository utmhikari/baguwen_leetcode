package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/16
 */
public class No_371_Sum_Of_Two_Integers {
    public int getSum(int a, int b) {
        int c;
        while (b != 0) {
            c = a & b;
            a ^= b;
            b = c << 1;
        }
        return a;
    }
}
