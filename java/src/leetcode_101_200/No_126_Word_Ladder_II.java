package leetcode_101_200;

import java.util.*;

/**
 * Created by 36191 on 2019/3/9
 */
public class No_126_Word_Ladder_II {
    
    List<List<String>> results;
    List<String> list;
    Map<String,List<String>> map;

    public List<List<String>> findLadders(String beginWord, String endWord, List<String> wordList) {
        Set<String> dict = new HashSet<>(wordList);
        return findLadders(beginWord, endWord, dict);
    }
    
    public List<List<String>> findLadders(String start, String end, Set<String> dict) {
        results = new ArrayList<>();
        if (dict.size() == 0 || !dict.contains(end)) { return results; }
        int curr = 1, next = 0;
        boolean found = false;
        list = new LinkedList<>();
        map = new HashMap<>();
        Queue<String> queue = new ArrayDeque<>();
        Set<String> unvisited = new HashSet<>(dict);
        Set<String> visited = new HashSet<>();
        queue.add(start);
        unvisited.add(end);
        unvisited.remove(start);
        while (!queue.isEmpty()) {
            String word = queue.poll();
            curr--;
            for (int i = 0; i < word.length(); i++){
                StringBuilder builder = new StringBuilder(word);
                for (char ch = 'a'; ch <= 'z'; ch++){
                    builder.setCharAt(i, ch);
                    String newWord = builder.toString();
                    if (unvisited.contains(newWord)){
                        if (visited.add(newWord)) { next++; queue.add(newWord); }
                        if (map.containsKey(newWord)) { map.get(newWord).add(word); }
                        else {
                            List<String> l = new LinkedList<>();
                            l.add(word);
                            map.put(newWord, l);
                        }
                        if (newWord.equals(end) && !found) { found=true; }
                    }
                }
            }
            if (curr == 0){
                if (found) { break; }
                curr = next; next = 0;
                unvisited.removeAll(visited);
                visited.clear();
            }
        }
        backTrace(end, start);
        return results;
    }

    private void backTrace(String word, String start){
        if (word.equals(start)){
            list.add(0,start);
            results.add(new ArrayList<>(list));
            list.remove(0);
            return;
        }
        list.add(0,word);
        if (map.get(word) != null) {
            for (String s : map.get(word)) {
                backTrace(s, start);
            }
        }
        list.remove(0);
    }

    public static void main(String[] args) {
        No_126_Word_Ladder_II solution = new No_126_Word_Ladder_II();
        String[] inputs = new String[] { "hot","dot","dog","lot","log","cog" };
        List<String> wl = new ArrayList<>(Arrays.asList(inputs));
        List<List<String>> ladders = solution.findLadders("hit", "cog", wl);
        for (List<String> ls : ladders) {
            for (String s : ls) { System.out.print(s + ", "); }
            System.out.println();
        }
    }
}
