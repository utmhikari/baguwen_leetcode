package leetcode_301_400;

import java.util.Stack;

/**
 * Created by 36191 on 2019/2/7
 */
public class No_372_Super_Pow {

    int pmod(int a, int x) {
        int mod = a % 1337, m = 1;
        for (int i = 0; i < x; i++) { m = m * mod % 1337; }
        return m;
    }

    int sp(int a, Stack<Integer> bDigits) {
        if (bDigits.empty()) { return 1; }
        int last = bDigits.pop();
        return pmod(sp(a, bDigits), 10) * pmod(a, last) % 1337;
    }

    public int superPow(int a, int[] b) {
        int mod = a % 1337;
        Stack<Integer> bDigits = new Stack<>();
        for (int i : b) { bDigits.push(i); }
        return sp(a, bDigits);
    }
}
