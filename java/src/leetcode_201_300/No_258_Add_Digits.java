package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/21
 */
public class No_258_Add_Digits {
    public int addDigits(int num) {
        return num == 0 ? 0 : ((num % 9 == 0) ? 9 : (num % 9));
    }
}
