package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/18
 * 066.加1
 */

public class No_066_Plus_One {
    public int[] plusOne(int[] digits) {
        int add = 0;
        int len = digits.length;
        digits[len - 1] = digits[len - 1] + 1;
        for(int i = len - 1; i >= 0; i--) {
            digits[i] = digits[i] + add;
            if(digits[i] > 9) {
                add = 1;
                digits[i] = digits[i] - 10;
            } else {
                add = 0;
            }
        }
        if(add == 1) {
            int[] newdigit = new int[len + 1];
            newdigit[0] = 1;
            System.arraycopy(digits, 0, newdigit, 1, len);
            return newdigit;
        } else {
            return digits;
        }
    }
}
