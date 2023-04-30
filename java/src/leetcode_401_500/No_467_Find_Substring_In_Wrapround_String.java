package leetcode_401_500;

import java.util.*;

public class No_467_Find_Substring_In_Wrapround_String {
    public int findSubstringInWraproundString(String p) {
        int len = p.length();
        if (len == 0) {
            return 0;
        }
        int[] charMaxCnts = new int[26];
        int startChar = -1, curCnt = 0;
        char[] charArray = p.toCharArray();
        for (int i = 0; i < len; i++) {
            char c = charArray[i];
            int curChar = (int)c - (int)'a';
            if (startChar == -1) {
                startChar = curChar;
                curCnt = 1;
                continue;
            }
            int expectedChar = (startChar + curCnt) % 26;
            if (curChar == expectedChar) {
                curCnt += 1;
            } else {
                for (int j = 0; j < curCnt; j++) {
                    int charIdx = (int)(charArray[i - j - 1]) - (int)'a';
                    charMaxCnts[charIdx] = Math.max(charMaxCnts[charIdx], j + 1);
                }
                startChar = curChar;
                curCnt = 1;
            }
        }
        for (int j = 0; j < curCnt; j++) {
            int charIdx = (int)(charArray[len - j - 1]) - (int)'a';
            charMaxCnts[charIdx] = Math.max(charMaxCnts[charIdx], j + 1);
        }
        int sum = 0;
        for (int i = 0; i < 26; i++) {
            if (charMaxCnts[i] != 0) {
                System.out.println(i + ", " + charMaxCnts[i]);
            }
            sum += charMaxCnts[i];
        }
        return sum;
    }

    public static void main(String[] args) {
        No_467_Find_Substring_In_Wrapround_String solution = new No_467_Find_Substring_In_Wrapround_String();
        String[] inputs = new String[] {
            "a",
            "cac",
            "zab",
            "zaba",
        };
        for (String input : inputs) {
            System.out.println(solution.findSubstringInWraproundString(input));
        }
    }

}
