package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/13
 */
public class No_342_Power_Of_Four {
    public boolean isPowerOfFour(int num) {
        if (num <= 0) { return false; }
        return num == 1 || ((num & (num - 1)) == 0 && (num % 10 == 4 || num % 10 == 6));
    }
}
