package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/27
 */

public class No_191_Number_of_1_bits {
    public int hammingWeight(int n) {
        int and = 1 << 31 - 1, w = 0;
        if (n < 0) { w++; n ^= Integer.MIN_VALUE; }
        while (n != 0) {
            if (n >= and) { n ^= and; w++;}
            and >>= 1;
        }
        return w;
    }

    public static void main(String[] args) {
        No_191_Number_of_1_bits solution = new No_191_Number_of_1_bits();
        System.out.println(solution.hammingWeight(11));
        System.out.println(solution.hammingWeight(128));
    }
}
