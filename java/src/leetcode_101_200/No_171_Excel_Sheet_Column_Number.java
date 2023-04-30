package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/25
 */

public class No_171_Excel_Sheet_Column_Number {
    public int titleToNumber(String s) {
        int len = s.length(), sum = 0;
        for (int i = 0; i < len; i++) { sum = sum * 26 + (s.charAt(i) - 64); }
        return sum;
    }

    public static void main(String[] args) {
        No_171_Excel_Sheet_Column_Number solution = new No_171_Excel_Sheet_Column_Number();
        String[] inputs = new String[] {"A", "AB", "ZY"};
        for (String s : inputs) {
            System.out.println(solution.titleToNumber(s));
        }
    }
}
