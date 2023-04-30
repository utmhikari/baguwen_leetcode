package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/10
 *
 */

public class No_036_Valid_Sudoku {

    public boolean isValidSudoku(char[][] board) {
        int[] nums = new int[9];
        for(int i = 0; i < 9; i+=3) {
            for(int j = 0; j < 9; j+=3) {
                for(int i1 = i; i1 < i + 3; i1++) {
                    for(int j1 = j; j1 < j + 3; j1++) {
                        char num = board[i1][j1];
                        int index = num - 49;
                        if (num != '.') {
                            if (nums[index] == 1) {
                                return false;
                            }
                            nums[index] = 1;
                        }
                    }
                }
                nums = new int[9];
            }
        }
        return true;
    }

    public void test() {
        char[][] testBoard = new char[][] {
                ".87654321".toCharArray(),
                "2........".toCharArray(),
                "3........".toCharArray(),
                "4........".toCharArray(),
                "5........".toCharArray(),
                "6........".toCharArray(),
                "7........".toCharArray(),
                "8........".toCharArray(),
                "9........".toCharArray()};
        System.out.println(isValidSudoku(testBoard));
    }
    public static void main(String[] args) {
        No_036_Valid_Sudoku solution = new No_036_Valid_Sudoku();
        solution.test();
    }
}
