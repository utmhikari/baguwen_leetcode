package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/17
 * 063. 路径数量（有障碍）
 */

public class No_063_Unique_Paths_II {

    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        int m = obstacleGrid.length;
        if(m == 0) {
            return 0;
        }
        int n = obstacleGrid[0].length;
        if(n == 0) {
            return 0;
        }
        int[][] grid = new int[m][n];
        if(obstacleGrid[0][0] == 1) {
            return 0;
        }
        grid[0][0] = 1;
        for(int i = 1; i < m; i++) {
            if(obstacleGrid[i][0] == 1) {
                grid[i][0] = 0;
            } else {
                grid[i][0] = grid[i - 1][0];
            }
        }
        for(int j = 1; j < n; j++) {
            if(obstacleGrid[0][j] == 1) {
                grid[0][j] = 0;
            } else {
                grid[0][j] = grid[0][j - 1];
            }
        }
        for(int i = 1; i < m; i++) {
            for(int j = 1; j < n; j++) {
                if(obstacleGrid[i][j] == 1) {
                    grid[i][j] = 0;
                } else {
                    grid[i][j] = grid[i - 1][j] + grid[i][j - 1];
                }
            }
        }
        return grid[m - 1][n - 1];
    }

    public void test() {
        int[][] obstacleGrid = new int[][] {
                new int[] {0, 0, 0},
                new int[] {0, 1, 0},
                new int[] {0, 0, 0}
        };
        System.out.println(uniquePathsWithObstacles(obstacleGrid));
    }

    public static void main(String[] args) {
        No_063_Unique_Paths_II solution = new No_063_Unique_Paths_II();
        solution.test();
    }
}
