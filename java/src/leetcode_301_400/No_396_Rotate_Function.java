package leetcode_301_400;

/**
 * Created by 36191 on 2019/2/12
 */
public class No_396_Rotate_Function {
    public int maxRotateFunction(int[] A) {
        int sum = 0, base = 0, len = A.length;
        for (int i = 0; i < len; i++) {
            sum += A[i]; base += i * A[i];
        }
        int max = base;
        for (int i = 0; i < len - 1; i++) {
            base += len * A[i] - sum;
            max = Math.max(base, max);
        }
        return max;
    }
}
