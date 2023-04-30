package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/25
 */
import java.util.*;

public class No_179_Largest_Number {


    public String largestNumber(int[] nums) {
        if (nums == null) { return "0"; }
        ArrayList<String> strs = new ArrayList<>();
        for (int i : nums) { strs.add(String.valueOf(i)); }
        strs.sort(new Comparator<String>() {
            @Override
            public int compare(String s1, String s2) {
                int l1 = s1.length(), l2 = s2.length();
                if (l1 == 0 && l2 != 0) { return -1; }
                else if (l2 == 0 && l1 != 0) { return 1; }
                else {
                    int i = 0;
                    while (i < l1 && i < l2) {
                        if (s1.charAt(i) > s2.charAt(i)) { return -1; }
                        else if (s1.charAt(i) < s2.charAt(i)) { return 1; }
                        else { ++i; }
                    }
                    if (i == l1 && i == l2) { return 0; }
                    else if (i == l1) { return compare(s1, s2.substring(i)); }
                    else { return compare(s1.substring(i), s2); }
                }
            }});
        if ("0".equals(strs.get(0))) { return "0"; }
        StringBuilder sb = new StringBuilder();
        for (String s : strs) { sb.append(s); }
        return sb.toString();
    }

    public static void main(String[] args) {
        No_179_Largest_Number solution = new No_179_Largest_Number();
        int[][] inputs = new int[][] {
                new int[] {10, 2},
                new int[] {3, 30, 34, 5, 9, 98},
                new int[] {3, 30, 34, 5, 9},
                new int[] {0, 0, 0},
                new int[] {824,938,1399,5607,6973,5703,9609,4398,8247}
        };

        for (int[] input : inputs) {
            System.out.println(solution.largestNumber(input));
        }
    }


}
