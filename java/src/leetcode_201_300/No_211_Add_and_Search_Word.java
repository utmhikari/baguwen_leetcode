package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/5
 */
public class No_211_Add_and_Search_Word {
    private class WordDictionary {

        WordDictionary[] children;
        boolean isEnd;

        /** Initialize your data structure here. */
        public WordDictionary() {
            this.children = new WordDictionary[26];
            this.isEnd = false;
        }

        /** Adds a word into the data structure. */
        public void addWord(String word) {
            int len = word.length();
            WordDictionary visit = this;
            for (int i = 0; i < len; i++) {
                int idx = word.charAt(i) - 'a';
                if (visit.children[idx] == null) { visit.children[idx] = new WordDictionary(); }
                visit = visit.children[idx];
            }
            visit.isEnd = true;
        }

        private boolean search(String word, int start, int len, WordDictionary visit) {
            for (int i = start; i < len; i++) {
                char c = word.charAt(i);
                if (c == '.') {
                    if (isEnd) { return true; }
                    for (int j = 0; j < 26; j++) {
                        if (visit.children[j] != null && search(word, i + 1, len, visit.children[j])) {
                            return true;
                        }
                    }
                    return false;
                } else {
                    int idx = word.charAt(i) - 'a';
                    if (visit.children[idx] == null) { return false; }
                    visit = visit.children[idx];
                }
            }
            return visit.isEnd;
        }

        /** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
        public boolean search(String word) {
            int len = word.length();
            if (len == 0) { return true; }
            WordDictionary visit = this;
            boolean isSearched = search(word, 0, len, visit);
            System.out.println(isSearched);
            return isSearched;
        }
    }

    public WordDictionary getWD() {
        return new WordDictionary();
    }

    public static void main(String[] args) {
        No_211_Add_and_Search_Word solution = new No_211_Add_and_Search_Word();
        WordDictionary wd = solution.getWD();
        wd.addWord("a");
        wd.addWord("a");
        wd.search(".");
        wd.search("a");
        wd.search("aa");
        wd.search("a");
        wd.search(".a");
        wd.search("a.");
    }
}
