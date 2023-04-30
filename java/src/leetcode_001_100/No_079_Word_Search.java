package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/21
 * 079. 寻找单词
 */

public class No_079_Word_Search {

    public boolean find(char[][] board, boolean[][] visitBoard, String word, int index, int r, int c) {
        if(index == word.length()) {
            return true;
        } else {
            if(index == 0) {
                for(int i = 0; i < board.length * board[0].length; i++) {
                    int row = i / board[0].length;
                    int col = i % board[0].length;
                    if(board[row][col] == word.charAt(index)) {
                        visitBoard[row][col] = true;
                        if(find(board, visitBoard, word, index + 1, row, col)) {
                            return true;
                        }
                    }
                    visitBoard[row][col] = false;
                }
                return false;
            } else {
                int[][] locations = new int[][] {
                        new int[] {r - 1, c},
                        new int[] {r, c - 1},
                        new int[] {r + 1, c},
                        new int[] {r, c + 1}
                };
                for(int[] loc : locations) {
                    if(loc[0] >= 0 && loc[0] < board.length && loc[1] >= 0 && loc[1] < board[0].length) {
                        if(!visitBoard[loc[0]][loc[1]] && board[loc[0]][loc[1]] == word.charAt(index)) {
                            visitBoard[loc[0]][loc[1]] = true;
                            if(find(board, visitBoard, word, index + 1, loc[0], loc[1])) {
                                return true;
                            }
                            visitBoard[loc[0]][loc[1]] = false;
                        }
                    }
                }
                return false;
            }
        }
    }

    public boolean exist(char[][] board, String word) {
        if(word.equals("")) {
            return true;
        }
        if(board.length == 0 || board[0].length == 0) {
            return false;
        }
        boolean[][] visitBoard = new boolean[board.length][board[0].length];
        return find(board, visitBoard, word, 0, -1, -1);
    }

    public void test() {
        char[][] board = new char[][] {
                new char[] {'a', 'b', 'c', 'e'},
                new char[] {'s', 'f', 'c', 's'},
                new char[] {'a', 'd', 'e', 'e'}
        };
        System.out.println(exist(board, "abcced"));
        System.out.println(exist(board, "see"));
        System.out.println(exist(board, "abcb"));
    }

    public static void main(String[] args) {
        No_079_Word_Search solution = new No_079_Word_Search();
        solution.test();
    }
}
