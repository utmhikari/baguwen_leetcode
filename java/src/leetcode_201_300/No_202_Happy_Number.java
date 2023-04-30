package leetcode_201_300;

/**
 * Created by 36191 on 2018/9/29
 * No.202 happy number 快乐数字
 */

import java.util.HashSet;

public class No_202_Happy_Number {

    private static HashSet<Integer> happy = new HashSet<Integer>() {{ add(1); }};
    private static HashSet<Integer> sad = new HashSet<Integer>() {{ add(0); }};

    private int cal(int n) {
        int sum = 0;
        while (n != 0) {
            int x = n % 10;
            sum += x * x;
            n /= 10;
        }
        return sum;
    }

    public boolean isHappy(int n) {
        if (happy.contains(n)) { return true; }
        if (sad.contains(n)) { return false; }
        HashSet<Integer> steps = new HashSet<>();
        steps.add(n);
        while (true) {
            n = cal(n);
            if (steps.contains(n) || sad.contains(n)) { sad.addAll(steps); return false; }
            if (happy.contains(n)) { happy.addAll(steps); return true; }
            steps.add(n);
        }
    }
}
