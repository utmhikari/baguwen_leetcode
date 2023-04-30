package leetcode_101_200;

import java.util.HashSet;

/**
 * Created by 36191 on 2018/10/4
 */
public class No_200_Number_of_Islands {

    private int encode(int x, int y, int n) {
        return x * n + y;
    }

    private int[] decode(int c, int n) {
        return new int[] { c / n, c % n };
    }

    private void search(int x, int y, int m, int n, boolean[][] visit, HashSet<Integer> ones, char[][] grid) {
        if (!ones.isEmpty() && (x >= 0 && x < m) && (y >= 0 && y < n) && !visit[x][y]) {
            visit[x][y] = true;
            if (grid[x][y] == '1') {
                ones.remove(encode(x, y, n));
                search(x - 1, y, m, n, visit, ones, grid);
                search(x + 1, y, m, n, visit, ones, grid);
                search(x, y - 1, m, n, visit, ones, grid);
                search(x, y + 1, m, n, visit, ones, grid);
            }
        }
    }

    public int numIslands(char[][] grid) {
        int m = grid.length;
        if (m == 0) { return 0; }
        else {
            int n = grid[0].length;
            if (n == 0) { return 0; }
            else {
                int cnt = 0;
                HashSet<Integer> ones = new HashSet<>();
                for (int i = 0; i < m; i++) {
                    for (int j = 0; j < n; j++) {
                        if (grid[i][j] == '1') { ones.add(encode(i, j, n));}
                    }
                }
                boolean[][] visit = new boolean[m][n];
                while (!ones.isEmpty()) {
                    int start = ones.iterator().next();
                    int[] coord = decode(start, n);
                    search(coord[0], coord[1], m, n, visit, ones, grid);
                    cnt++;
                }
                return cnt;
            }
        }
    }

    public static void main(String[] args) {
        No_200_Number_of_Islands solution = new No_200_Number_of_Islands();
        char[][] grid = new char[][] {
                new char[] {'0', '1', '0', '1', '0'},
                new char[] {'1', '1', '0', '1', '0'},
                new char[] {'0', '1', '1', '1', '0'},
                new char[] {'0', '0', '1', '0', '1'}
        };
        System.out.println(solution.numIslands(grid));
    }
}
