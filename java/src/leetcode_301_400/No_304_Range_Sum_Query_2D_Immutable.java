package leetcode_301_400;

/**
 * Created by 36191 on 2018/11/17
 */
public class No_304_Range_Sum_Query_2D_Immutable {
    private int[][] m;

    public void NumMatrix(int[][] matrix) {
        if (matrix.length != 0 && matrix[0].length != 0) {
            m = new int[matrix.length][matrix[0].length];
            for (int i = 0; i < matrix.length; i++) {
                for (int j = 0; j < matrix[i].length; j++) {
                    if (j == 0) { m[i][j] = matrix[i][j]; }
                    else { m[i][j] = matrix[i][j] + m[i][j - 1]; }
                }
            }
        }
    }

    public int sumRegion(int row1, int col1, int row2, int col2) {
        int sum = 0;
        for (int i = row1; i <= row2; i++) {
            sum += m[i][col2];
            if (col1 != 0) { sum -= m[i][col1 - 1]; }
        }
        return sum;
    }
}
