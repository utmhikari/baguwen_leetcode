package leetcode_101_200;

/**
 * Created by 36191 on 2018/8/9
 */

import java.util.*;

public class No_130_Surrounded_Regions {
    private static HashSet<Integer> positions = new HashSet<>();

    private int _code(int x, int y, int h) {
        return x * h + y;
    }

    private int[] _decode(int c, int h) {
        int x = c / h;
        int y = c % h;
        return new int[]{x, y};
    }

    private void findO(char[][] board, int start, int l, int h) {
        int[] xy = _decode(start, h);
        int x = xy[0], y = xy[1];
        if(!positions.contains(start)) {
            positions.add(start);
            System.out.println(start);
            if(x != 0 && board[x - 1][y] == 'O') {
                findO(board, _code(x-1, y, h), l, h);
            }
            if(x != l - 1 && board[x + 1][y] == 'O') {
                findO(board, _code(x+1, y, h), l, h);
            }
            if(y != 0 && board[x][y - 1] == 'O') {
                findO(board, _code(x, y-1, h), l, h);
            }
            if(y != h - 1 && board[x][y + 1] == 'O') {
                findO(board, _code(x, y+1, h), l, h);
            }
        }
    }

    public void solve(char[][] board) {
        if(board.length != 0 && board[0].length != 0) {
            if(!positions.isEmpty()) {positions.clear();}
            int l = board.length;
            int h = board[0].length;
            System.out.println(l + " --- " + h);
            ArrayList<Integer> starts = new ArrayList<>();
            for(int y = 0; y < h; y++) {
                if(board[0][y] == 'O'){ starts.add(_code(0, y, h)); }
                if(board[l - 1][y] == 'O'){ starts.add(_code(l - 1, y, h));}
            }
            for(int x = 1; x < l - 1; x++) {
                if(board[x][0] == 'O'){ starts.add(_code(x, 0, h));}
                if(board[x][h - 1] == 'O'){ starts.add(_code(x, h - 1, h));}
            }
            for(int s : starts) {
                findO(board, s, l, h);
            }
            for(int x = 0; x < l; x++) {
                for(int y = 0; y < h; y++) {
                    if(board[x][y] == 'O' && !positions.contains(_code(x, y, h))) { board[x][y] = 'X';}
                }
            }
        }
    }
}
