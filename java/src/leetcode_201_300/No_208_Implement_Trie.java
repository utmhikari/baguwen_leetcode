package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/4
 */
public class No_208_Implement_Trie {
    private class Trie {

        Trie visit;
        Trie[] children;
        boolean isEnd;

        /** Initialize your data structure here. */
        public Trie() {
            this.children = new Trie[26];
            this.isEnd = false;
        }

        /** Inserts a word into the trie. */
        public void insert(String word) {
            int len = word.length();
            visit = this;
            for (int i = 0; i < len; i++) {
                int idx = word.charAt(i) - 'a';
                if (visit.children[idx] == null) { visit.children[idx] = new Trie(); }
                visit = visit.children[idx];
            }
            visit.isEnd = true;
        }

        /** Returns if the word is in the trie. */
        public boolean search(String word) {
            return startsWith(word) && visit.isEnd;
        }

        /** Returns if there is any word in the trie that starts with the given prefix. */
        public boolean startsWith(String prefix) {
            int len = prefix.length();
            visit = this;
            for (int i = 0; i < len; i++) {
                int idx = prefix.charAt(i) - 'a';
                if (visit.children[idx] == null) { return false; }
                visit = visit.children[idx];
            }
            return true;
        }
    }

    public Trie trie() {
        return new Trie();
    }

    public static void main(String[] args) {
        No_208_Implement_Trie solution = new No_208_Implement_Trie();
        Trie trie = solution.trie();
        trie.insert("apple");
        System.out.println(trie.search("apple"));
        System.out.println(trie.search("app"));
        System.out.println(trie.startsWith("app"));
        trie.insert("app");
        System.out.println(trie.search("app"));
    }
}
