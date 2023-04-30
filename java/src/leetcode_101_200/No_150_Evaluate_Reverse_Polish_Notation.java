package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/22
 */
import java.util.*;

public class No_150_Evaluate_Reverse_Polish_Notation {

    public int operate(char op, int v1, int v2) {
        if (op == '+') { return v1 + v2; }
        else if (op == '-') { return v1 - v2; }
        else if (op == '*') { return v1 * v2; }
        else { return v1 / v2; }
    }

    // returns the eval value and the next rIdx
    public int[] eval(String[] tokens, int rIdx, HashSet<Character> ops) {
        char c = tokens[rIdx].charAt(0);
        if (ops.contains(c) && tokens[rIdx].length() == 1) {
            int[] firstReturns = eval(tokens, rIdx - 1, ops);
            int[] secondReturns = eval(tokens, firstReturns[1], ops);
            return new int[] {operate(c, secondReturns[0], firstReturns[0]), secondReturns[1]};
        } else {
            return new int[] { Integer.parseInt(tokens[rIdx]), rIdx - 1};
        }
    }

    public int evalRPN(String[] tokens) {
        int len = tokens.length;
        HashSet<Character> operators = new HashSet<Character>() {{
            add('+'); add('-'); add('*'); add('/');
        }};
        return eval(tokens, len - 1, operators)[0];
    }

    public static void main(String[] args) {
        No_150_Evaluate_Reverse_Polish_Notation solution = new No_150_Evaluate_Reverse_Polish_Notation();
        ArrayList<String[]> ts = new ArrayList<String[]>() {{
            add(new String[] {"2", "1", "+", "3", "*"});
            add(new String[] {"4", "13", "5", "/", "+"});
            add(new String[] {"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"});
        }};
        for (String[] t : ts) {
            System.out.println(solution.evalRPN(t));
        }
    }
}
