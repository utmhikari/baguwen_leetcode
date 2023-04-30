package leetcode_301_400;

import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

/**
 * Created by 36191 on 2019/2/7
 */
public class No_394_Decode_String {
    public String decodeString(String s) {
        Stack<List<Character>> strStack = new Stack<>();
        strStack.push(new ArrayList<Character>());
        Stack<Integer> intStack = new Stack<>();
        int i = 0, l = s.length();
        while (i < l) {
            char c = s.charAt(i);
            if (c <= '9' && c >= '0') {
                int times = c - '0';
                ++i;
                while (s.charAt(i) <= '9' && s.charAt(i) >= '0') {
                    times = times * 10 + (s.charAt(i++) - '0');
                }
                strStack.push(new ArrayList<Character>());
                intStack.push(times);
            } else if (c == ']') {
                int times = intStack.pop();
                List<Character> headStr = strStack.pop();
                for (int j = 0; j < times; j++) {
                    for (int k = 0; k < headStr.size(); k++) {
                        strStack.peek().add(headStr.get(k));
                    }
                }
            } else {
                strStack.peek().add(c);
            }
            ++i;
        }
        StringBuilder sb = new StringBuilder();
        List<Character> finalStr = strStack.pop();
        for (int x = 0; x < finalStr.size(); x++) { sb.append(finalStr.get(x)); }
        return sb.toString();
    }
}
