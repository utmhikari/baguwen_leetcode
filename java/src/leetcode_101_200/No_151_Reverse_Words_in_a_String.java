package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/24
 */

import java.util.*;

public class No_151_Reverse_Words_in_a_String {
    public String reverseWords(String s) {
        String[] splits = s.trim().split(" ");
        int len = splits.length;
        if (len == 0) { return ""; }
        StringBuilder sb = new StringBuilder();
        for (int i = len - 1; i >= 0; i--) {
            sb.append(splits[i]);
            if (splits[i].length() != 0) { sb.append(" "); }
        }
        return sb.toString().trim();
    }

    public static void main(String[] args) {
        No_151_Reverse_Words_in_a_String solution = new No_151_Reverse_Words_in_a_String();
        String s1 = "the sky is blue";
        System.out.println(solution.reverseWords(s1));
        String s2 = "   a   b ";
        System.out.println(solution.reverseWords(s2));
    }
}
