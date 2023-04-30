package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/23
 */
public class No_289_Game_Of_Life {

    private boolean checkValid(int[][] board, int x, int y) {
        return x >= 0 && x < board.length && y >= 0 && y < board[x].length;
    }

    public void gameOfLife(int[][] board) {
        int[][] nextState = new int[board.length][board[0].length];
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[i].length; j++) {
                int cnt = 0;
                for (int x = i - 1; x <= i + 1; x++) {
                    for (int y = j - 1; y <= j + 1; y++) {
                        if (checkValid(board, x, y) && !(x == i && y == j)) {
                            if (board[x][y] == 1) { cnt++; }
                        }
                    }
                }
                if (cnt == 3 || (board[i][j] == 1 && cnt == 2)) { nextState[i][j] = 1; }
            }
        }
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[i].length; j++) {
                board[i][j] = nextState[i][j];
            }
        }
    }
}
