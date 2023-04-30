package leetcode_401_500;

/**
 * Created by 36191 on 2019/7/7
 */
public class No_413_Arithmetic_Slices {

    public int numberOfArithmeticSlices(int[] A) {
        int len = A.length;
        if (len < 3) {
            return 0;
        }
        int num = 0, cnt = 0, last = A[1] - A[0];
        for (int i = 2; i < A.length; ++i) {
            int diff = A[i] - A[i - 1];
            if (diff != last) {
                num += cnt * (cnt + 1) / 2;
                cnt = 0;
                last = diff;
            } else {
                ++cnt;
            }
        }
        return num + cnt * (cnt + 1) / 2;
    }

    public static void main(String[] args) {
        No_413_Arithmetic_Slices solution = new No_413_Arithmetic_Slices();
    }
}
