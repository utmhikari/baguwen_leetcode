package leetcode_301_400;

/**
 * Created by 36191 on 2018/11/17
 */

public class No_306_Additive_Number {

    private boolean find(String num, long first, long second, int start) {
        if (start >= num.length()) { return false; }
        long sum = first + second;
        String sumStr = String.valueOf(sum);
        String remain = num.substring(start);
        if (remain.startsWith(sumStr)) {
            if (remain.length() == sumStr.length()) { return true; }
            else { return find(num, second, sum, start + sumStr.length()); }
        }
        return false;
    }

    public boolean isAdditiveNumber(String num) {
        System.out.println("\nNum: " + num);
        int len = num.length(), mid = len / 2;
        if (len < 3) { return false; }
        int firstIdx = 0;
        while (num.charAt(firstIdx) == '0') {
            ++firstIdx;
            if (firstIdx == len) { break; }
        }
        if (firstIdx == len) { return true; }
        else if (firstIdx == 0) {
            for (int i = 0; i < mid; i++) {
                long first = Long.parseLong(num.substring(0, i + 1));
                if (num.charAt(i + 1) == '0') {
                    if (find(num, first, 0, i + 2)) { return true; }
                } else {
                    int secMid = (len + i + 1) / 2;
                    for (int j = i + 1; j < secMid; j++) {
                        long second = Long.parseLong(num.substring(i + 1, j + 1));
                        if (find(num, first, second, j + 1)) { return true; }
                    }
                }
            }
        } else if (firstIdx == 1) {
            for (int i = 1; i < mid; i++) {
                long second = Long.parseLong(num.substring(1, i + 1));
                if (find(num, 0, second, i + 1)) { return true; }
            }
        }
        return false;
    }

    public static void main(String[] args) {
        No_306_Additive_Number solution = new No_306_Additive_Number();
        System.out.println(solution.isAdditiveNumber("112358"));
        System.out.println(solution.isAdditiveNumber("199100199"));
        System.out.println(solution.isAdditiveNumber("1991001998"));
        System.out.println(solution.isAdditiveNumber("101"));
        System.out.println(solution.isAdditiveNumber("0235813"));
        System.out.println(solution.isAdditiveNumber("0"));
        System.out.println(solution.isAdditiveNumber("0000000"));
        System.out.println(solution.isAdditiveNumber("011235"));
    }
}
