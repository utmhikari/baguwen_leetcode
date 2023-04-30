package leetcode_201_300;

/**
 * Created by 36191 on 2018/9/29
 * No.204 count primes
 */

public class No_204_Count_Primes {


    private void gen(boolean[] vis, int n) {
        int m = (int)Math.round(Math.sqrt((double)n + 0.5));
        for (int i = 2; i <= m; i++) {
            if (!vis[i - 1]) {
                for (int j = i * i; j <= n; j += i) {
                    vis[j - 1] = true;
                }
            }
        }
    }

    public int countPrimes(int n) {
        boolean[] vis = new boolean[n];
        gen(vis, n);
        int sum = 0;
        for (int i = 2; i < n; i++) {
            if (!vis[i - 1]) { sum += 1; }
        }
        return sum;
    }

    public static void main(String[] args) {
        No_204_Count_Primes solution = new No_204_Count_Primes();
        System.out.println(solution.countPrimes(10));
    }
}
