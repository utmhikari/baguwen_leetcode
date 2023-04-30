package leetcode_301_400;


/**
 * Created by 36191 on 2018/11/28
 */
public class No_313_Super_Ugly_Number {

    private int minFactors(int[] factors) {
        int min = Integer.MAX_VALUE;
        for (int i : factors) {
            if (i < min) { min = i; }
        }
        return min;
    }
    public int nthSuperUglyNumber(int n, int[] primes) {
        if (n == 1) { return 1; }
        int len = primes.length;
        int[] factors = new int[len], indices = new int[len], uglyNums = new int[n];
        System.arraycopy(primes, 0, factors, 0, len);
        for (int i = 1; i < n; i++) {
            int min = minFactors(factors);
            uglyNums[i] = min;
            for (int k = 0; k < factors.length; k++) {
                if (min == factors[k]) {
                    factors[k] = primes[k] * uglyNums[++indices[k]];
                }
            }
        }
        return uglyNums[n - 1];
    }

    public static void main(String[] args) {
        No_313_Super_Ugly_Number solution = new No_313_Super_Ugly_Number();
        System.out.println(solution.nthSuperUglyNumber(12, new int[] {2, 7, 13, 19}));
    }
}
