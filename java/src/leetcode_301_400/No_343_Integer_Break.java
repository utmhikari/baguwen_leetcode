package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/13
 */
public class No_343_Integer_Break {
    public int integerBreak(int n) {
        double max = 1, digit;
        if (n == 2) { return 1; }
        int half = n / 2 + 1;
        for (int i = 2; i <= half; i++) {
            int div = n / i;
            int mod = n - i * div;
            if (mod == 0) {
                digit = Math.pow(i, div);
            } else {
                digit = Math.pow(i, div - 1);
                if (div == 1) {
                    digit = mod * (n - mod);
                } else if (mod == 1) {
                    digit *= (mod + i);
                } else {
                    digit *= (mod * i);
                }
            }
            max = Math.max(max, digit);
        }
        return (int)max;
    }
    
    public static void main(String[] args) {
        No_343_Integer_Break solution = new No_343_Integer_Break();
        for (int i = 2; i <= 58; i++) {
            System.out.println(i + ": " + solution.integerBreak(i));
        }
    }
    
}
