package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/4
 */
public class No_201_Bitwise_AND_of_Numbers_Range {

    private static int sum;

    private void cal(long m, long n) {
        for (int i = 31; i >= 0; i--) {
            long left = (long)1 << i;
            if (m >= left) {
                if (n < (left << 1)) {
                    sum += left;
                    cal(m - left, n - left);
                } else { break; }
            }
        }
    }

    public int rangeBitwiseAnd(int m, int n) {
        sum = 0;
        cal((long)m, (long)n);
        return sum;
    }

    public static void main(String[] args) {
        No_201_Bitwise_AND_of_Numbers_Range solution = new No_201_Bitwise_AND_of_Numbers_Range();
        int[][] inputs = new int[][] {
                new int[] {5, 7},
                new int[] {0, 1},
                new int[] {12, 15},
        };
        for (int[] input : inputs) { System.out.println(solution.rangeBitwiseAnd(input[0], input[1])); }
    }
}
