package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/16
 * 051. N皇后问题
 */
import java.util.List;
import java.util.LinkedList;
public class No_051_N_Queens {

    public void addAnswer(List<List<String>> answers, int cols[], int n) {
        LinkedList<String> ans = new LinkedList<>();
        for(int row = 0; row < n; row++) {
            StringBuilder sb = new StringBuilder();
            for(int col = 0; col < n; col++) {
                if(col == cols[row]) {
                    sb.append('Q');
                } else {
                    sb.append('.');
                }
            }
            ans.add(sb.toString());
        }
        answers.add(ans);
    }

    public boolean isValid(int[] cols, int row, int col) {
        for(int _row = 0; _row < row; _row++) {
            int _col = cols[_row];
            if(_col == col || (_row + _col == row + col) || (_row - _col == row - col)) {
                return false;
            }
        }
        return true;
    }

    public void solve(List<List<String>> answers, int cols[], int row, int n) {
        for(int col = 0; col < n; col++) {
            if(isValid(cols, row, col)) {
                cols[row] = col;
                if(row == n - 1) {
                    addAnswer(answers, cols, n);
                } else {
                    solve(answers, cols, row + 1, n);
                }
            }
        }
    }

    public List<List<String>> solveNQueens(int n) {
        List<List<String>> answers = new LinkedList<>();
        if(n <= 0) {
            return answers;
        }
        int cols[] = new int[n];
        for(int i = 0; i < n; i++) {
            cols[i] = -1;
        }
        int row = 0;
        solve(answers, cols, row, n);
        return answers;
    }

    public void test(int n) {
        List<List<String>> answers = solveNQueens(n);
        for(List<String> ans : answers) {
            for(String row : ans) {
                System.out.println(row);
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        No_051_N_Queens solution = new No_051_N_Queens();
        solution.test(4);
        solution.test(5);
        solution.test(8);
    }
}
