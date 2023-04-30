package leetcode_001_100;

import java.util.Stack;

/**
 * Created by 36191 on 2017/10/10.
 * 032.最长括号对
 */
public class No_032_Longest_Valid_Parentheses {

    /*
    public int longestValidParentheses(String s) {
        int n = s.length(), longest = 0;
        Stack<Integer> st = new Stack<>();
        for (int i = 0; i < n; i++) {
            if (s.charAt(i) == '(') {
                st.push(i);
            }
            else {
                if (!st.empty()) {
                    if (s.charAt(st.peek()) == '(') {
                        st.pop();
                    }
                    else {
                        st.push(i);
                    }
                }
                else {
                    st.push(i);
                }
            }
        }
        if (st.empty()) {
            longest = n;
        } else {
            int a = n, b = 0;
            while (!st.empty()) {
                b = st.pop();
                longest = longest > a-b-1 ? longest : a-b-1;
                a = b;
            }
            longest = longest > a ? longest : a;
        }
        return longest;
    }


     */

    public boolean isValid(String s, int startIndex, int endIndex) {
        Stack<Character> brackets = new Stack<>();
        for (int i = startIndex; i <= endIndex; i++) {
            char c = s.charAt(i);
            switch (c) {
                case '(':
                    brackets.push(c);
                    break;
                case ')':
                    if (brackets.empty() || brackets.peek() != '(') {
                        return false;
                    } else {
                        brackets.pop();
                    }
                    break;
                default:
                    break;
            }
        }
        return brackets.empty();
    }

    public String pre(String s) {
        StringBuilder sb = new StringBuilder();
        int l = s.length();
        for(int i = 0; i < l; i++) {
            sb.append(s.charAt(i));
            sb.append('#');
        }
        sb.deleteCharAt(sb.length() - 1);
        return sb.toString();
    }

    public int longestValidParentheses(String s) {
        if(s.equals("")) {
            return 0;
        }
        String preStr = pre(s);
        int maxLen = 1;
        int len = preStr.length();
        for(int i = 0; i < len; i++) {
            if(preStr.charAt(i) == '#') {
                int parLen = 1;
                int k = 1;
                while(true) {
                    int startIndex = i - k;
                    int endIndex = i + k;
                    if(startIndex < 0 || endIndex > len) {
                        if(parLen > maxLen) {
                            maxLen = parLen;
                        }
                        break;
                    }
                    int tempLen = 1 + 2 * k;
                    if(isValid(preStr, startIndex, endIndex)) {
                        parLen = tempLen;
                    }
                    k += 2;
                }
            }
        }
        return (maxLen == 1) ? 0 : ((maxLen + 1) / 2);
    }

    public void output(String s) {
        System.out.println(longestValidParentheses(s));
    }

    public static void main(String[] args) {
        No_032_Longest_Valid_Parentheses solution = new No_032_Longest_Valid_Parentheses();
        solution.output(")(");
        solution.output("(()");
        solution.output(")()())");
        solution.output("(()())");
    }
}
