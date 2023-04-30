package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/17
 * 058. 最后单词的长度
 */

public class No_058_Length_Of_Last_Word {
    public int lengthOfLastWord(String s) {
        int lastLen = 0;
        int temp = 0;
        int len = s.length();
        for(int i = 0; i < len; i++) {
            if(s.charAt(i) == ' ') {
                if(temp != 0) {
                    lastLen = temp;
                    temp = 0;
                }
            } else {
                temp++;
            }
        }
        if(temp != 0) {
            lastLen = temp;
            temp = 0;
        }
        return lastLen;
    }

    public static void main(String[] args) {
        No_058_Length_Of_Last_Word solution = new No_058_Length_Of_Last_Word();
        System.out.println(solution.lengthOfLastWord("b   a    "));
    }
}
