package leetcode_101_200;

/**
 * Created by 36191 on 2018/8/9
 */

import java.util.*;

public class No_125_Valid_Palindrome {
    public boolean isPalindrome(String s) {
        if(s.length() < 2) {
            return true;
        }
        s = s.toLowerCase();
        ArrayList<Character> chars = new ArrayList<>();
        for(int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if(c >= '0' && c <= '9' || c >= 'a' && c <= 'z') {
                chars.add(c);
            }
        }
        int l = 0, r = chars.size() - 1;
        while(l < r) {
            if(chars.get(l++) != chars.get(r--)) {
                return false;
            }
        }
        return true;
    }

}
