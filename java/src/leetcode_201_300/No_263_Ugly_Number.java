package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/21
 */
public class No_263_Ugly_Number {
    public boolean isUgly(int num) {
        if (num < 1) { return false; }
        int[] primes = new int[] {2, 3, 5};
        for (int i = 0; i < 3; i++) {
            while (num % primes[i] == 0) { num /= primes[i]; }
        }
        return num == 1;
    }
}
