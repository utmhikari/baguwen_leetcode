package leetcode_201_300;

import java.util.*;

/**
 * Created by 36191 on 2019/8/18
 */
public class No_212_Word_Search_II {
    public class Tree {
        private HashMap<Character, Tree> map;
        private boolean end;

        public Tree() {
            map = new HashMap<>();
            end = false;
        }

        public boolean hasChar(char c) {
            return map.containsKey(c);
        }

        public void putChar(char c) {
            map.put(c, new Tree());
        }

        public Tree getTree(char c) {
            return hasChar(c) ? map.get(c) : null;
        }

        public void setEnd(boolean b) {
            end = b;
        }

        public boolean isEnd() {
            return end;
        }

        public void insert(String s) {
            Tree root = this;
            for (int i = 0, l = s.length(); i < l; ++i) {
                char c = s.charAt(i);
                if (!root.hasChar(c)) {
                    root.putChar(c);
                }
                root = root.getTree(c);
            }
            root.setEnd(true);
        }
    }

    private void dfs(char[][] board, Tree tree, HashSet<String> set, HashSet<Integer> path, StringBuffer sb, int h, int w, int y, int x) {
        int code = y * w + x;
        if (y < 0 || y >= h || x < 0 || x >= w || path.contains(code)) {
            return;
        }
        char c = board[y][x];
        if (tree.hasChar(c)) {
            sb.append(c);
            path.add(code);
            Tree next = tree.getTree(c);
            if (next.isEnd()) {
                set.add(sb.toString());
            }
            dfs(board, next, set, path, sb, h, w, y - 1, x);
            dfs(board, next, set, path, sb, h, w, y + 1, x);
            dfs(board, next, set, path, sb, h, w, y, x - 1);
            dfs(board, next, set, path, sb, h, w, y, x + 1);
            sb.deleteCharAt(sb.length() - 1);
            path.remove(code);
        }
    }

    public List<String> findWords(char[][] board, String[] words) {
        List<String> ans = new ArrayList<>();
        int h = board.length;
        if (h == 0) {
            return ans;
        }
        int w = board[0].length;
        if (w == 0) {
            return ans;
        }
        HashSet<String> set = new HashSet<>();
        HashSet<Integer> path = new HashSet<>();
        Tree tree = new Tree();
        StringBuffer sb = new StringBuffer();
        for (String s : words) {
            tree.insert(s);
        }
        for (int y = 0; y < h; ++y) {
            for (int x = 0; x < w; ++x) {
                dfs(board, tree, set, path, sb, h, w, y, x);
            }
        }
        if (set.size() > 0) {
            ans.addAll(set);
            Collections.sort(ans);
        }
        return ans;
    }
}
