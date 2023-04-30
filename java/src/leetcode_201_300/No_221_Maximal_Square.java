package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/7
 */
public class No_221_Maximal_Square {
    public int maximalSquare(char[][] matrix) {
        if (matrix.length == 0 || matrix[0].length == 0) { return 0; }
        int max = 0, x = matrix.length, y = matrix[0].length;
        int[][][] records = new int[x][y][3];
        for (int i = 0; i < x; i++) {
            for (int j = 0; j < y; j++) {
                if (matrix[i][j] == '0') { records[i][j] = new int[]{0, 0, 0}; }
                else {
                    records[i][j][0] = i == 0 ? 1 : records[i - 1][j][0] + 1;
                    records[i][j][1] = j == 0 ? 1 : records[i][j - 1][1] + 1;
                    records[i][j][2] = (i == 0 || j == 0) ? 1 : Math.min(records[i - 1][j - 1][2] + 1,
                            Math.min(records[i][j][0], records[i][j][1]));
                    max = Math.max(max, records[i][j][2]);
                }
            }
        }
        return max * max;
    }
}
