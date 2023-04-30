package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/6.
 * 017.电话号码字母组合
 */

import java.util.List;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Arrays;

public class No_017_Letter_Combination_Phone_Number {

    public static HashMap<Integer, List<String>> letterMap = new HashMap<Integer, List<String>>() {{
        put(2, Arrays.asList("a", "b", "c"));
        put(3, Arrays.asList("d", "e", "f"));
        put(4, Arrays.asList("g", "h", "i"));
        put(5, Arrays.asList("j", "k", "l"));
        put(6, Arrays.asList("m", "n", "o"));
        put(7, Arrays.asList("p", "q", "r", "s"));
        put(8, Arrays.asList("t", "u", "v"));
        put(9, Arrays.asList("w", "x", "y", "z"));
    }};

    public List<String> combineCombList(List<String> lb1, List<String> lb2) {
        List<String> combineList = new ArrayList<>();
        for(String s1 : lb1) {
            for(String s2 : lb2) {
                combineList.add(s1 + s2);
            }
        }
        return combineList;
    }

    public List<String> letterCombinations(String digits) {
        int len = digits.length();
        if(len == 0) {
            return new ArrayList<String>();
        } else if(len == 1) {
            int num = Integer.parseInt(digits);
            if(letterMap.containsKey(num)) {
                return letterMap.get(num);
            } else {
                return new ArrayList<String>();
            }
        } else if(len == 2) {
            return combineCombList(letterCombinations(Character.toString(digits.charAt(0))),
                    letterCombinations(Character.toString(digits.charAt(1))));
        } else {
            return combineCombList(letterCombinations(digits.substring(0, len / 2)),
                    letterCombinations(digits.substring(len / 2, len)));
        }
    }

    public void output(String digits) {
        for(String s : letterCombinations(digits)) {
            System.out.print(s + ",");
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_017_Letter_Combination_Phone_Number solution = new No_017_Letter_Combination_Phone_Number();
        solution.output("22");
        solution.output("23");
    }
}
