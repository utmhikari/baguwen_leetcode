package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/16
 * No.127 单词阶梯
 */

import java.util.*;

public class No_127_Word_Ladder {

    private boolean dist(String s1, String s2) {
        int d = 0;
        for (int i = 0; i < s1.length(); i++) {
            if (s1.charAt(i) != s2.charAt(i)) { d++; }
            if (d > 1) { return false; }
        }
        return d == 1;
    }

    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        if (!wordList.contains(endWord)) { return 0; }
        int len = wordList.size();
        wordList.add(beginWord);
        boolean[][] graph = new boolean[len + 1][len + 1];
        for (int i = 0; i < len; i++) {
            for (int j = i + 1; j <= len; j++) {
                graph[i][j] = dist(wordList.get(i), wordList.get(j));
            }
        }
        int dist = 0;
        int eIdx = wordList.indexOf(endWord);
        HashSet<Integer> all = new HashSet<>();
        HashSet<Integer> prev;
        HashSet<Integer> temp = new HashSet<>();
        temp.add(len);
        all.add(len);
        while (!temp.isEmpty()) {
            prev = new HashSet<>(temp);
            temp.clear();
            dist += 1;
            for (int i : prev) {
                for (int j = 0; j < len; j++) {
                    if ((graph[i][j] || graph[j][i]) && !all.contains(j)) {
                        if (j == eIdx) { return dist + 1; }
                        else { temp.add(j); all.add(j); }
                    }
                }
            }
        }
        return 0;
    }

    public static void main(String[] args) {
        No_127_Word_Ladder solution = new No_127_Word_Ladder();
        List<String> wordList = new ArrayList<String>() {{
           add("hot"); add("dot"); add("dog"); add("lot"); add("log"); add("cog");
        }};
        System.out.println(solution.ladderLength("hit", "cog", wordList));
    }
}
