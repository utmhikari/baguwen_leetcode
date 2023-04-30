package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/26
 */

public class No_190_Reverse_Bits {

    private long[] mi = new long[32];

    public int reverseBits(int n) {
        if (mi[0] == 0) { for (int i = 0; i < 32; i++) { mi[i] = (long)1 << i; } }
        long sum = 0;
        for (int i = 31; i >= 0; i--) {
            if ((long)n >= mi[i]) {
                System.out.print(i + ", " + mi[i]);
                sum += mi[31 - i];
                n -= mi[i];
            }
        }
        System.out.println("\n" + Integer.toBinaryString((int)sum));
        return (int)sum;
    }

//    public int reverseBits(int n) {
//        int result = 0;
//
//        if (n == 0)
//            return result;
//
//        for (int i = 0; i < 32; i++) {
//            result <<= 1;
//            result += (n & 1);
//
//            n >>= 1;
//        }
//
//        return result;
//    }

    public static void main(String[] args) {
        No_190_Reverse_Bits solution = new No_190_Reverse_Bits();
        System.out.println(solution.reverseBits(1));
    }
}
