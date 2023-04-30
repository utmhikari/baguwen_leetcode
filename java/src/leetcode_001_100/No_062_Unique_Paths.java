package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/17
 * 062.路径数量
 */

public class No_062_Unique_Paths {

    public int uniquePaths(int m, int n) {
        if(m == 0 || n == 0) {
            return 0;
        } else {
            int[][] grid = new int[m][n];
            grid[0][0] = 1;
            for(int i = 1; i < m; i++) {
                grid[i][0] = 1;
            }
            for(int j = 1; j < n; j++) {
                grid[0][j] = 1;
            }
            for(int i = 1; i < m; i++) {
                for(int j = 1; j < n; j++) {
                    grid[i][j] = grid[i - 1][j] + grid[i][j - 1];
                }
            }
            return grid[m - 1][n - 1];
        }
    }

    public static void main(String[] args) {
        No_062_Unique_Paths solution = new No_062_Unique_Paths();
    }
}
