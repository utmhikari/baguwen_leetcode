package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/15
 * 048.旋转图像
 */

public class No_048_Rotate_Image {

    public void rotateIndex(int[][] matrix, int i1, int i2, int i3, int i4) {
        int temp = matrix[i1][i2];
        matrix[i1][i2] = matrix[i4][i1];
        matrix[i4][i1] = matrix[i3][i4];
        matrix[i3][i4] = matrix[i2][i3];
        matrix[i2][i3] = temp;
    }

    public void rotateOutside(int[][] matrix, int startIndex, int endIndex, int matrixLen) {
        for(int i = startIndex; i < endIndex; i++) {
            int index1 = startIndex;
            int index2 = i;
            int index3 = matrixLen - 1 - startIndex;
            int index4 = matrixLen - 1 - i;
            rotateIndex(matrix, index1, index2, index3, index4);
        }
    }

    public void rotate(int[][] matrix) {
        int len = matrix.length;
        if(matrix.length <= 1) {
            return;
        }
        for(int i = 0; i < len / 2; i++) {
            rotateOutside(matrix, i, len - 1 - i, len);
        }
        for(int i = 0; i < len; i++) {
            for(int j = 0; j < len; j++) {
                System.out.print(matrix[i][j] + ", ");
            }
            System.out.println();
        }
        System.out.println();
    }

    public void test() {
        int[][] matrix1 = new int[][] {
                new int[] {1, 2, 3},
                new int[] {4, 5, 6},
                new int[] {7, 8, 9}
        };
        rotate(matrix1);
        int[][] matrix2 = new int[][] {
                new int[] {5, 1, 9, 11},
                new int[] {2, 4, 8, 10},
                new int[] {13, 3, 6, 7},
                new int[] {15, 14, 12, 16}
        };
        rotate(matrix2);
        int[][] matrix3 = new int[][] {
                new int[] {1, 2, 3, 4, 5, 6},
                new int[] {7, 8, 9, 10, 11, 12},
                new int[] {13, 14, 15, 16, 17, 18},
                new int[] {19, 20, 21, 22, 23, 24},
                new int[] {25, 26, 27, 28, 29, 30},
                new int[] {31, 32, 33, 34, 35, 36}
        };
        rotate(matrix3);
    }

    public static void main(String[] args) {
        No_048_Rotate_Image solution = new No_048_Rotate_Image();
        solution.test();
    }
}
