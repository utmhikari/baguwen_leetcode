package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/16
 */
public class No_367_Valid_Perfect_Square {
    public boolean isPerfectSquare(int num) {
        for (int i = 1; ; i++) {
            int j = num / i;
            if (j == i && num % i == 0) { return true; }
            else if (j < i || (j == i && num % i != 0)) { break; }
        }
        return false;
    }
}
