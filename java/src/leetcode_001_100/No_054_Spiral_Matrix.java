package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/16
 * 054. 旋转矩阵
 */
import java.util.List;
import java.util.LinkedList;

public class No_054_Spiral_Matrix {

    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> order = new LinkedList<>();
        int row = matrix.length;
        if(row == 0) {
            return order;
        }
        int col = matrix[0].length;
        if(col == 0) {
            return order;
        }
        int count = 0;
        int all = row * col;
        int r = 0;
        int c = 0;
        int nextR = 0;
        int nextC = col - 1;
        while(count < all) {
            order.add(matrix[r][c]);
            if(r == nextR) {
                if(c < nextC) {
                    c++;
                } else if(c > nextC) {
                    c--;
                } else {
                    if(nextR < row / 2) {
                        nextR = row - 1 - nextR;
                        r++;
                    } else {
                        nextR = row - nextR;
                        r--;
                    }
                    if(nextC < col / 2) {
                        nextC = col - 2 - nextC;
                    } else {
                        nextC = col - 1 - nextC;
                    }
                }
            } else if(r < nextR) {
                r++;
            } else {
                r--;
            }
            count++;
        }
        return order;
    }

    public void test() {
        int[][] matrix1 = new int[][] {
                new int[] {1, 2, 3, 4, 5},
                new int[] {6, 7, 8, 9, 10},
                new int[] {11, 12, 13, 14, 15},
                new int[] {16, 17, 18, 19, 20},
        };
        List<Integer> order1 = spiralOrder(matrix1);
        for(int i : order1) {
            System.out.print(i + ", ");
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_054_Spiral_Matrix solution = new No_054_Spiral_Matrix();
        solution.test();
    }
}
