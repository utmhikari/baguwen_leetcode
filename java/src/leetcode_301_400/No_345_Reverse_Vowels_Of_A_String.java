package leetcode_301_400;

import java.util.ArrayList;
import java.util.HashSet;

/**
 * Created by 36191 on 2019/1/13
 */
public class No_345_Reverse_Vowels_Of_A_String {
    public String reverseVowels(String s) {
        if (s.length() == 0) { return ""; }
        HashSet<Character> vowels = new HashSet<Character>() {{
            add('a'); add('e'); add('i'); add('o'); add('u');
            add('A'); add('E'); add('I'); add('O'); add('U');
        }};
        char[] arr = s.toCharArray();
        ArrayList<Integer> indices = new ArrayList<>();
        for (int i = 0; i < arr.length; i++) {
            if (vowels.contains(arr[i])) { indices.add(i); }
        }
        if (indices.size() == 0) { return s; }
        for (int i = 0, l = indices.size() - 1; i <= l / 2; i++) {
            char tmp = arr[indices.get(i)];
            arr[indices.get(i)] = arr[indices.get(l - i)];
            arr[indices.get(l - i)] = tmp;
        }
        return new String(arr);
    }
}
