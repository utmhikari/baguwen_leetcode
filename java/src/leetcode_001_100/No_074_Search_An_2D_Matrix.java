package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/20
 * 074. 矩阵中搜索值
 */

public class No_074_Search_An_2D_Matrix {

    public boolean searchInRow(int[][] matrix, int startCol, int endCol, int row, int target) {
        if(startCol == endCol) {
            return matrix[row][startCol] == target;
        }
        int halfCol = startCol + (endCol - startCol) / 2;
        if(target > matrix[row][halfCol]) {
            return searchInRow(matrix, halfCol + 1, endCol, row, target);
        } else {
            return searchInRow(matrix, startCol, halfCol, row, target);
        }
    }

    public int confirmRow(int[][] matrix, int startRow, int endRow, int col, int target) {
        if(startRow == endRow) {
            if(target >= matrix[startRow][0] && target <= matrix[startRow][col - 1]) {
                return startRow;
            } else {
                return -1;
            }
        }
        int halfRow = startRow + (endRow - startRow) / 2;
        if(target > matrix[halfRow][0]) {
            if(target <= matrix[halfRow][col - 1]) {
                return halfRow;
            } else {
                return confirmRow(matrix, halfRow + 1, endRow, col, target);
            }
        } else {
            return confirmRow(matrix, startRow, halfRow, col, target);
        }
    }

    public boolean searchMatrix(int[][] matrix, int target) {
        int m = matrix.length;
        if(m == 0) {
            return false;
        }
        int n = matrix[0].length;
        if(n == 0) {
            return false;
        }
        int row = confirmRow(matrix, 0, m - 1, n, target);
        if(row == -1) {
            return false;
        }
        return searchInRow(matrix, 0, n - 1, row, target);
    }

    public void test() {
        int[][] matrix = new int[][] {
                new int[] {1, 3, 5, 7},
                new int[] {10, 11, 16, 20},
                new int[] {23, 30, 34, 50}
        };
        System.out.println(searchMatrix(matrix, 3));
        System.out.println(searchMatrix(matrix, 17));
        System.out.println(searchMatrix(matrix, 34));
    }

    public static void main(String[] args) {
        No_074_Search_An_2D_Matrix solution = new No_074_Search_An_2D_Matrix();
        solution.test();
    }
}
