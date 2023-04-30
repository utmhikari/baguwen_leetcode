package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/17
 * 059. 生成旋转矩阵
 */

public class No_059_Spiral_Matrix_II {
    public int[][] generateMatrix(int n) {
        int[][] matrix = new int[n][n];
        int count = 1;
        int all = n * n;
        int r = 0;
        int c = 0;
        int nextR = 0;
        int nextC = n - 1;
        while(count <= all) {
            matrix[r][c] = count++;
            if(r == nextR) {
                if(c < nextC) {
                    c++;
                } else if(c > nextC) {
                    c--;
                } else {
                    if(nextR < n / 2) {
                        nextR = n - 1 - nextR;
                        r++;
                    } else {
                        nextR = n - nextR;
                        r--;
                    }
                    if(nextC < n / 2) {
                        nextC = n - 2 - nextC;
                    } else {
                        nextC = n - 1 - nextC;
                    }
                }
            } else if(r < nextR) {
                r++;
            } else {
                r--;
            }
        }
        return matrix;
    }
}
