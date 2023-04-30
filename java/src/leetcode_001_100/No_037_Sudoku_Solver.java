package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/11
 * 037.解数独
 */

public class No_037_Sudoku_Solver {

    // 上两倍、左两倍、下两倍、右两倍、上差1、左差1、右差1、下差1
    private static int[][] rules;

    private void init() {
        rules = new int[9][9];
        rules[0][0] = 0b00100000; rules[0][3] = 0b00000010; rules[0][7] = 0b00000001; rules[0][8] = 0b00000001;
        rules[1][1] = 0b00000001; rules[1][3] = 0b00000001; rules[1][6] = 0b00000101; rules[1][8] = 0b00001001;
        rules[2][0] = 0b00000001; rules[2][2] = 0b00000001; rules[2][3] = 0b00011001; rules[2][4] = 0b01000001;
        rules[2][5] = 0b00000001; rules[2][6] = 0b00001000; rules[2][8] = 0b00001000;
        rules[3][0] = 0b00001010; rules[3][1] = 0b00000100; rules[3][2] = 0b00001000; rules[3][3] = 0b00001000;
        rules[3][4] = 0b00001000; rules[3][5] = 0b00001000; rules[3][6] = 0b00000010; rules[3][7] = 0b00000100;
        rules[4][1] = 0b00000011; rules[4][2] = 0b00000100; rules[4][3] = 0b00000011; rules[4][4] = 0b00010100;
        rules[4][5] = 0b01000010; rules[4][6] = 0b00000100; rules[4][7] = 0b00000010; rules[4][8] = 0b00000100;
        rules[5][0] = 0b00000010; rules[5][1] = 0b00001100; rules[5][2] = 0b00000010; rules[5][3] = 0b00001100;
        rules[5][5] = 0b00000010; rules[5][6] = 0b00000110; rules[5][7] = 0b00000100; rules[5][8] = 0b00000001;
        rules[6][0] = 0b00100000; rules[6][4] = 0b00000010; rules[6][5] = 0b00000100; rules[6][6] = 0b00100000;
        rules[6][7] = 0b00000010; rules[6][8] = 0b00001100;
        rules[7][0] = 0b10000000; rules[7][1] = 0b00000001; rules[7][2] = 0b00000010; rules[7][3] = 0b00000101;
        rules[7][4] = 0b00000001; rules[7][6] = 0b10000001; rules[7][7] = 0b00100000;
        rules[8][1] = 0b00001000; rules[8][3] = 0b00001000; rules[8][4] = 0b00011000; rules[8][5] = 0b01000000;
        rules[8][6] = 0b00011000; rules[8][7] = 0b11000000;
    }

    private boolean checkRule(char[][] board, int x, int y, char ans) {
        if (rules[x][y] != 0) {
            int a = ans - '0';
            for (int i = 7; i >= 0; --i) {
                int flag = (1 << i);
                if ((rules[x][y] & flag) == flag) {
                    int b;
                    if (i > 3) {
                        if (i == 7) { b = board[x - 1][y]; }
                        else if (i == 6) { b = board[x][y - 1]; }
                        else if (i == 5) { b = board[x + 1][y]; }
                        else { b = board[x][y + 1]; }
                        if (b != '.') { b -= '0'; if (b * 2 != a && a * 2 != b) { return false; } }
                    } else {
                        if (i == 3) { b = board[x - 1][y]; }
                        else if (i == 2) { b = board[x][y - 1]; }
                        else if (i == 1) { b = board[x][y + 1]; }
                        else { b = board[x + 1][y]; }
                        if (b != '.') { b -= '0'; if (b - 1 != a && b + 1 != a) { return false; } }
                    }
                }
            }
        }
        return true;
    }

    public boolean isValid(char[][] board, int row, int col, char ans) {
        for(int x = 0; x < 9; x++) {
            if(board[x][col] == ans && x != row) {
                return false;
            }
            if(board[row][x] == ans && x != col) {
                return false;
            }
        }
        int startRow = (row / 3) * 3;
        int startCol = (col / 3) * 3;
        for(int x = startRow; x < startRow + 3; x++) {
            for(int y = startCol; y < startCol + 3; y++) {
                if(board[x][y] == ans && !(x == row && y == col)) {
                    return false;
                }
            }
        }
        return checkRule(board, row, col, ans);
    }

    public boolean solve(char[][] board) {
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                if (board[i][j] == '.') {
                    for(char c = '1'; c <= '9'; c++) {
                        if(isValid(board, i, j, c)) {
                            board[i][j] = c;
                            if (solve(board)) {
                                return true;
                            } else {
                                board[i][j] = '.';
                            }
                        }
                    }
                    return false;
                }
            }
        }
        return true;
    }


    public void solveSudoku(char[][] board) {
        solve(board);
        outputSolution(board);
    }


    public void outputSolution(char[][] board) {
        System.out.println();
        for(int i = 0; i < 9; i++) {
            for(int j = 0; j < 9; j++) {
                System.out.print(board[i][j]);
                System.out.print(", ");
            }
            System.out.println();
        }
    }

    public void test() {
        init();
        char[][] testBoard = new char[][] {
                "......1..".toCharArray(),
                "........3".toCharArray(),
                ".........".toCharArray(),
                "........1".toCharArray(),
                ".........".toCharArray(),
                ".........".toCharArray(),
                "......6..".toCharArray(),
                "......31.".toCharArray(),
                "......42.".toCharArray()};
        solveSudoku(testBoard);
    }
    public static void main(String[] args) {
        No_037_Sudoku_Solver solution = new No_037_Sudoku_Solver();
        solution.test();
    }
}
