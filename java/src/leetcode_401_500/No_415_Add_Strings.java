package leetcode_401_500;

import java.util.Stack;

/**
 * Created by 36191 on 2019/2/12
 */
public class No_415_Add_Strings {

    public String addStrings(String num1, String num2) {
        Stack<Integer> chars = new Stack<>();
        int inc = 0, i = 0, len1 = num1.length(), len2 = num2.length();
        while (true) {
            int idx1 = len1 - 1 - i, idx2 = len2 - 1 - i;
            if (idx1 < 0 && idx2 < 0) { if (inc > 0) { chars.push(inc); } break; }
            int n1 = idx1 < 0 ? 0 : num1.charAt(idx1) - '0', n2 = idx2 < 0 ? 0 : num2.charAt(idx2) - '0';
            int sum = n1 + n2 + inc;
            chars.push(sum % 10);
            inc = sum / 10;
            ++i;
        }
        StringBuilder sb = new StringBuilder();
        while (!chars.empty()) { sb.append(chars.pop()); }
        return sb.toString();
    }

    public static void main(String args) {

    }
}
