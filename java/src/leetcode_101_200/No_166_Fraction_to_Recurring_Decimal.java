package leetcode_101_200;

import java.util.ArrayList;
import java.util.HashMap;

/**
 * Created by 36191 on 2018/10/3
 * No.166 除法小数
 */
public class No_166_Fraction_to_Recurring_Decimal {

    private String encode(long quo, long mod) {
        return String.valueOf(quo) + "_" + String.valueOf(mod);
    }

    public String fractionToDecimal(int numerator, int denominator) {
        if (denominator == 0) { return ""; }
        StringBuilder dec = new StringBuilder();
        long left = (long)numerator / (long)denominator, n = (long)numerator % (long)denominator, d = (long)denominator;
        dec.append(String.valueOf(left));
        if (n != 0) {
            if (left == 0 && ((numerator > 0 && denominator < 0) || (numerator < 0 && denominator > 0))) {
                dec.insert(0, '-');
            }
            dec.append('.');
            ArrayList<Character> digits = new ArrayList<>();
            HashMap<String, Integer> map = new HashMap<>(32);
            while (true) {
                n *= 10;
                long quo = n / d, mod = n % d;
                if (mod == 0) { digits.add((char)(Math.abs(quo) + '0')); break; }
                else {
                    String code = encode(quo, mod);
                    if (map.containsKey(code)) {
                        digits.add(map.get(code), '(');
                        digits.add(')');
                        break;
                    } else {
                        digits.add((char)(Math.abs(quo) + '0'));
                        map.put(code, digits.size() - 1);
                        n = mod;
                    }
                }
            }
            for (char c : digits) { dec.append(c); }
        }
        return dec.toString();
    }

    public static void main(String[] args) {
        No_166_Fraction_to_Recurring_Decimal solution = new No_166_Fraction_to_Recurring_Decimal();
        int[][] inputs = new int[][] {
                new int[] {1, 2},
                new int[] {2, 1},
                new int[] {2, 3},
                new int[] {13, 7},
                new int[] {1, 6},
                new int[] {50, -8},
                new int[] {1, -6},
                new int[] {1, Integer.MIN_VALUE},
                new int[] {Integer.MIN_VALUE, 1}
        };
        for (int[] input : inputs) {
            System.out.println(solution.fractionToDecimal(input[0], input[1]));
        }
    }
}
