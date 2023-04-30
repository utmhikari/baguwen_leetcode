package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/17
 * 064.最短路径长度
 */

public class No_064_Minimum_Path_Sum {

    public int minPathSum(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int[][] minPathSums = new int[m][n];
        minPathSums[0][0] = grid[0][0];
        for(int i = 1; i < m; i++) {
            minPathSums[i][0] = minPathSums[i - 1][0] + grid[i][0];
        }
        for(int j = 1; j < n; j++) {
            minPathSums[0][j] = minPathSums[0][j - 1] + grid[0][j];
        }
        for(int i = 1; i < m; i++) {
            for(int j = 1; j < n; j++) {
                minPathSums[i][j] = Math.min(minPathSums[i - 1][j], minPathSums[i][j - 1]) + grid[i][j];
            }
        }
        return minPathSums[m - 1][n - 1];
    }

    public static void main(String[] args) {
        No_064_Minimum_Path_Sum solution = new No_064_Minimum_Path_Sum();
    }
}
