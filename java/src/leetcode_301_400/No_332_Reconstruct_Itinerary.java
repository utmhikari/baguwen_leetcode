package leetcode_301_400;

import java.util.*;

/**
 * Created by 36191 on 2019/1/10
 */
public class No_332_Reconstruct_Itinerary {
    public List<String> findItinerary(String[][] tickets) {
        Map<String, LinkedList<String>> graph = new HashMap<>(32);
        for (String[] fromTo : tickets){
            if (!graph.containsKey(fromTo[0])) {
                graph.put(fromTo[0], new LinkedList<>());
            }
            graph.get(fromTo[0]).add(fromTo[1]);
        }
        for (String key: graph.keySet()){
            Collections.sort(graph.get(key));
        }
        LinkedList<String> path = new LinkedList<>();
        dfs("JFK", path, graph);
        return path;
    }
    
    private void dfs(String node, LinkedList<String> path, Map<String, LinkedList<String>> graph){
        if (graph.containsKey(node)) {
            while (graph.get(node).size() != 0) {
                String nei = graph.get(node).poll();
                dfs(nei, path, graph);
            }
        }
        path.addFirst(node);
    }
    
    public static void main(String[] args) {
        No_332_Reconstruct_Itinerary solution = new No_332_Reconstruct_Itinerary();
        String[][][] inputs = new String[][][] {
                new String[][] {
                        new String[] { "MUC", "LHR" },
                        new String[] { "JFK", "MUC" },
                        new String[] { "SFO", "SJC" },
                        new String[] { "LHR", "SFO" },
                },
                new String[][] {
                        new String[] { "JFK", "SFO" },
                        new String[] { "JFK", "ATL" },
                        new String[] { "SFO", "ATL" },
                        new String[] { "ATL", "JFK" },
                        new String[] { "ATL", "SFO" },
                },
                new String[][] {
                        new String[] { "JFK", "KUL" },
                        new String[] { "JFK", "NRT" },
                        new String[] { "NRT", "JFK" },
                },
        };
        for (String[][] input : inputs) {
            List<String> ans = solution.findItinerary(input);
            System.out.print("The answer is: ");
            for (String s : ans) {
                System.out.print(s + ", ");
            }
            System.out.println();
        }
    }
}
