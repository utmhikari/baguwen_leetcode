package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/16
 * 052. N皇后问题（答案数量）
 */

public class No_052_N_Queens_II {

    public boolean isValid(int[] cols, int row, int col) {
        for(int _row = 0; _row < row; _row++) {
            int _col = cols[_row];
            if(_col == col || (_row + _col == row + col) || (_row - _col == row - col)) {
                return false;
            }
        }
        return true;
    }

    public void solve(int[] answers, int cols[], int row, int n) {
        for(int col = 0; col < n; col++) {
            if(isValid(cols, row, col)) {
                cols[row] = col;
                if(row == n - 1) {
                    answers[n - 1]++;
                } else {
                    solve(answers, cols, row + 1, n);
                }
            }
        }
    }

    public int totalNQueens(int n) {
        int[] answers = new int[20];
        if(n <= 0) {
            return 0;
        }
        int cols[] = new int[n];
        for(int i = 0; i < n; i++) {
            cols[i] = -1;
        }
        int row = 0;
        solve(answers, cols, row, n);
        return answers[n - 1];
    }

    public void test(int n) {
        System.out.println(totalNQueens(n));
    }

    public static void main(String[] args) {
        No_052_N_Queens_II solution = new No_052_N_Queens_II();
        solution.test(1);
        solution.test(2);
        solution.test(3);
        solution.test(4);
        solution.test(5);
        solution.test(6);
        solution.test(7);
        solution.test(8);
    }


}
