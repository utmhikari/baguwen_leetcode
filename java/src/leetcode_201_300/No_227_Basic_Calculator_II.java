package leetcode_201_300;

import java.util.ArrayList;

/**
 * Created by 36191 on 2018/10/14
 */

public class No_227_Basic_Calculator_II {

    private int cal(String expr) {
        int digit = 0, val = 0, mulOrDiv = -1;
        for (int i = 0, l = expr.length(); i < l; i++) {
            char c = expr.charAt(i);
            if (c == '*' || c == '/') {
                if (mulOrDiv == -1) { val = digit; }
                else if (mulOrDiv == 0) { val *= digit; }
                else { val /= digit; }
                digit = 0;
                mulOrDiv = (c == '*') ? 0 : 1;
            } else { digit = digit * 10 + (c - '0'); }
        }
        if (mulOrDiv == -1) { val = digit; }
        else if (mulOrDiv == 0) { val *= digit; }
        else { val /= digit; }
        return val;
    }

    public int calculate(String s) {
        ArrayList<String> exprs = new ArrayList<>();
        ArrayList<Boolean> plusOrMinusList = new ArrayList<>();
        boolean plusOrMinus = true;
        StringBuilder sb = new StringBuilder();
        for (int i = 0, l = s.length(); i < l; i++) {
            char c = s.charAt(i);
            if (c != ' ') {
                if (c == '+' || c == '-') {
                    exprs.add(sb.toString());
                    sb.delete(0, sb.length());
                    plusOrMinusList.add(plusOrMinus);
                    plusOrMinus = (c == '+');
                } else { sb.append(c); }
            }
        }
        exprs.add(sb.toString());
        plusOrMinusList.add(plusOrMinus);
        int ans = 0;
        if (exprs.size() == plusOrMinusList.size()) {
            for (int i = 0, l = exprs.size(); i < l; i++) {
                ans += plusOrMinusList.get(i) ? cal(exprs.get(i)) : -cal(exprs.get(i));
            }
        }
        return ans;
    }
}
