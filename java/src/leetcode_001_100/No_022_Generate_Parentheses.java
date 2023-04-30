package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/7.
 * 022.生成括号组合
 */
import java.util.List;
import java.util.ArrayList;
import java.util.HashSet;

public class No_022_Generate_Parentheses {

    public List<String> generateParenthesis(int n) {
        List<String> parentheses = new ArrayList<>();
        if(n <= 0) {
            return parentheses;
        } else if(n == 1) {
            parentheses.add("()");
            return parentheses;
        } else {
            HashSet<String> parenthesesSet = new HashSet<>();
            int half = n / 2;
            for(int i = 1; i <= half; i++) {
                List<String> l1 = generateParenthesis(i);
                List<String> l2 = generateParenthesis(n - i);
                if(i == 1) {
                    for(String s : l2) {
                        parenthesesSet.add("(" + s + ")");
                        parenthesesSet.add("()" + s);
                        parenthesesSet.add(s + "()");
                    }
                } else {
                    for(String s1 : l1) {
                        for(String s2 : l2) {
                            parenthesesSet.add(s1 + s2);
                            if(n - i != i) {
                                parenthesesSet.add(s2 + s1);
                            }
                        }
                    }
                }
            }
            parentheses.addAll(parenthesesSet);
            return parentheses;
        }
    }

    public void output(int n) {
        for(String s : generateParenthesis(n)) {
            System.out.println(s);
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_022_Generate_Parentheses solution = new No_022_Generate_Parentheses();
        solution.output(0);
        solution.output(1);
        solution.output(2);
        solution.output(3);
        solution.output(4);
    }
}
