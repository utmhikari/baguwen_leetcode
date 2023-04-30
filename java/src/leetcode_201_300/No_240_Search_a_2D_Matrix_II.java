package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/20
 */

public class No_240_Search_a_2D_Matrix_II {

    private boolean isValid(int[][] matrix, boolean[][] visit, int h, int w) {
        return !(h < 0 || h >= matrix.length || w < 0 || w >= matrix[0].length || visit[h][w]);
    }

    private boolean search(int[][] matrix, boolean[][] visit, int target, int h, int w) {
        if (!isValid(matrix, visit, h, w)) { return false; }
        if (matrix[h][w] == target) { return true; }
        else {
            visit[h][w] = true;
            if (matrix[h][w] > target) {
                return search(matrix, visit, target, h - 1, w) || search(matrix, visit, target, h, w - 1);
            } else {
                return search(matrix, visit, target, h + 1, w) || search(matrix, visit, target, h, w + 1);
            }
        }
    }

    public boolean searchMatrix(int[][] matrix, int target) {
        if (matrix.length == 0 || matrix[0].length == 0) { return false; }
        boolean[][] visit = new boolean[matrix.length][matrix[0].length];
        return search(matrix, visit, target, matrix.length / 2, matrix[0].length / 2);
    }
}
