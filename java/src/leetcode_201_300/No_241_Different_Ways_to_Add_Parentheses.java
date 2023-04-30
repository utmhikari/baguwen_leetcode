package leetcode_201_300;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by 36191 on 2018/10/20
 */

public class No_241_Different_Ways_to_Add_Parentheses {

    public List<Integer> diffWaysToCompute(String input) {
        List<Integer> ans = new ArrayList<>();
        boolean hasSym = false;
        for (int i = 0; i < input.length(); i++) {
            char c = input.charAt(i);
            if (!(c >= '0' && c <= '9')) {
                hasSym = true;
                List<Integer> left = diffWaysToCompute(input.substring(0, i));
                List<Integer> right = diffWaysToCompute(input.substring(i + 1));
                for (int x : left) {
                    for (int y : right) {
                        if (c == '*') { ans.add(x * y); }
                        else if (c == '-') { ans.add(x - y); }
                        else { ans.add(x + y); }
                    }
                }
            }
        }
        if (!hasSym) { ans.add(Integer.parseInt(input)); }
        return ans;
    }

    public static void main(String[] args) {
        No_241_Different_Ways_to_Add_Parentheses solution = new No_241_Different_Ways_to_Add_Parentheses();
        String[] inputs = new String[] {
                "2-1-1",
                "2*3-4*5",
        };
        for (String s : inputs) {
            for (int i : solution.diffWaysToCompute(s)) {
                System.out.print(i + ", ");
            }
            System.out.println();
        }
    }
}
