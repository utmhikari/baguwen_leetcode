package leetcode_401_500;

import java.util.LinkedList;

/**
 * Created by 36191 on 2019/6/2
 */
public class No_402_Remove_K_Digits {
    public String removeKdigits(String num, int k) {
        int len = num.length();
        if (len == k) {
            return "0";
        }
        LinkedList<Character> list = new LinkedList<>();
        int i = 0;
        while (i < len) {
            char c = num.charAt(i);
            while (k > 0 && !list.isEmpty() && list.getLast() > c) {
                list.removeLast();
                k--;
            }
            list.addLast(c);
            ++i;
        }
        while (k > 0) {
            list.removeLast();
            k--;
        }
        while (list.size() > 1 && list.getFirst() == '0') {
            list.removeFirst();
        }
        StringBuilder sb = new StringBuilder();
        while (!list.isEmpty()) {
            sb.append(list.removeFirst());
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        No_402_Remove_K_Digits solution = new No_402_Remove_K_Digits();
        System.out.println(solution.removeKdigits("1432219", 3));
    }
}

