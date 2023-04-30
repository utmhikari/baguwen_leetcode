package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/7.
 * 020.括号匹配
 */
import java.util.Stack;
public class No_020_Valid_Parentheses {

    public boolean isValid(String s) {
        if (s == null) {
            return false;
        } else if (s.equals("")) {
            return true;
        } else {
            Stack<Character> brackets = new Stack<>();
            int len = s.length();
            for(int i = 0; i < len; i++) {
                char c = s.charAt(i);
                switch(c) {
                    case '(':
                        brackets.push(c);
                        break;
                    case '[':
                        brackets.push(c);
                        break;
                    case '{':
                        brackets.push(c);
                        break;
                    case ')':
                        if(brackets.empty() || brackets.peek() != '(') {
                            return false;
                        } else {
                            brackets.pop();
                        }
                        break;
                    case ']':
                        if(brackets.empty() || brackets.peek() != '[') {
                            return false;
                        } else {
                            brackets.pop();
                        }
                        break;
                    case '}':
                        if(brackets.empty() || brackets.peek() != '{') {
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
    }

    public static void main(String[] args) {
        No_020_Valid_Parentheses solution = new No_020_Valid_Parentheses();
        System.out.println(solution.isValid(""));
        System.out.println(solution.isValid("()"));
        System.out.println(solution.isValid("([])"));
        System.out.println(solution.isValid("([]){(([])[])}"));
        System.out.println(solution.isValid("()[]{}"));
        System.out.println(solution.isValid("([)]"));
    }
}
