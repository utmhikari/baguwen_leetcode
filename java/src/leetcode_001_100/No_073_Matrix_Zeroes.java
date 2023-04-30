package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/19
 * 073. 设置矩阵0值
 */

public class No_073_Matrix_Zeroes {

    /*
    // Put the rowZero and colZero sign to the first row/col, then back and consider the first row!
    // So that we can only use constant space!
    // https://leetcode.com/problems/set-matrix-zeroes/description/
    public void setZeroes(int[][] matrix) {
        boolean fr = false,fc = false;
        for(int i = 0; i < matrix.length; i++) {
            for(int j = 0; j < matrix[0].length; j++) {
                if(matrix[i][j] == 0) {
                    if(i == 0) fr = true;
                    if(j == 0) fc = true;
                    matrix[0][j] = 0;
                    matrix[i][0] = 0;
                }
            }
        }
        for(int i = 1; i < matrix.length; i++) {
            for(int j = 1; j < matrix[0].length; j++) {
                if(matrix[i][0] == 0 || matrix[0][j] == 0) {
                    matrix[i][j] = 0;
                }
            }
        }
        if(fr) {
            for(int j = 0; j < matrix[0].length; j++) {
                matrix[0][j] = 0;
            }
        }
        if(fc) {
            for(int i = 0; i < matrix.length; i++) {
                matrix[i][0] = 0;
            }
        }
    }
     */

    public void setZeroes(int[][] matrix) {
        int m = matrix.length;
        int n = matrix[0].length;
        boolean[] cols = new boolean[n];
        for(int i = 0; i < m; i++) {
            boolean rowZero = false;
            for(int j = 0; j < n; j++) {
                if(matrix[i][j] == 0) {
                    rowZero = true;
                    cols[j] = true;
                }
            }
            if(rowZero) {
                for(int j = 0; j < n; j++) {
                    matrix[i][j] = 0;
                }
            }
        }
        for(int j = 0; j < n; j++) {
            if(cols[j]) {
                for(int i = 0; i < m; i++) {
                    matrix[i][j] = 0;
                }
            }
        }
    }

    public static void main(String[] args) {
        No_073_Matrix_Zeroes solution = new No_073_Matrix_Zeroes();
    }
}
