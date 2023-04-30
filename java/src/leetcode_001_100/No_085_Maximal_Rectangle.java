package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/24
 * 085. 最大矩形面积
 */

public class No_085_Maximal_Rectangle {



    public int maximalRectangle(char[][] matrix) {
        int area = 0;
        int h = matrix.length;
        if (h == 0) {
            return area;
        }
        int w = matrix[0].length;
        if (w == 0) {
            return area;
        }
        int[][] right = new int[h][w];
        for (int i = 0; i < h; ++i) {
            for (int j = 0; j < w; ++j) {
                if (matrix[i][j] == '1') {
                    right[i][j] = j == 0 ? 1 : right[i][j - 1] + 1;
                    int height, width = right[i][j];
                    for (height = i; height >= 0 && matrix[height][j] != '0'; --height) {
                        width = Math.min(right[height][j], width);
                        area = Math.max(area, width * (i - height + 1));
                    }
                }
            }
        }
        return area;
    }

    public static void main(String[] args) {
        No_085_Maximal_Rectangle solution = new No_085_Maximal_Rectangle();
    }
}
