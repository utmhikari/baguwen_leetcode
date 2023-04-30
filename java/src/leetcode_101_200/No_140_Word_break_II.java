package leetcode_101_200;

import java.util.*;

/**
 * Created by 36191 on 2019/4/21
 */
public class No_140_Word_break_II {

    private List<String> dfs(String s, HashSet<String> dict, HashMap<String, LinkedList<String>> map) {
        if (map.containsKey(s)) {
            return map.get(s);
        }
        LinkedList<String> res = new LinkedList<>();
        if (s.length() == 0) {
            res.add("");
            return res;
        }
        for (String word : dict) {
            if (s.startsWith(word)) {
                List<String> sublist = dfs(s.substring(word.length()), dict, map);
                for (String sub : sublist) {
                    res.add(word + (sub.isEmpty() ? "" : " ") + sub);
                }
            }
        }
        map.put(s, res);
        return res;
    }

    public List<String> wordBreak(String s, List<String> wordDict) {
        HashSet<String> dict = new HashSet<>(wordDict);
        return dfs(s, dict, new HashMap<>(32));
    }

    public static void main(String[] args) {
        No_140_Word_break_II solution = new No_140_Word_break_II();
        List<String> wb1 = solution.wordBreak("catsanddog", new ArrayList<String>() {{
            add("aaaa"); add("aaa");
        }});
        for (String s : wb1) {
            System.out.println(s);
        }
    }
}
