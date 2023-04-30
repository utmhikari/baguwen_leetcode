package leetcode_301_400;


import java.util.*;

/**
 * Created by 36191 on 2019/1/26
 */
public class No_399_Evaluate_Division {

    private HashSet<String> zeros;
    private HashSet<String> visited;
    private HashMap<String, HashMap<String, Double>> nodes;
    
    private void addToNodes(String[][] equations, double[] values) {
        int len = equations.length;
        for (int i = 0; i < len; i++) {
            String[] eq = equations[i];
            double val = values[i];
            if (!nodes.containsKey(eq[0])) {
                nodes.put(eq[0], new HashMap<>(16));
            }
            nodes.get(eq[0]).put(eq[1], val);
            if (val != 0.0) {
                if (!nodes.containsKey(eq[1])) {
                    nodes.put(eq[1], new HashMap<>(16));
                }
                nodes.get(eq[1]).put(eq[0], 1.0 / val);
            } else {
                zeros.add(eq[0]);
            }
        }
    }

    /**
     * initialization: non zero, all contained (but maybe in different graph), not equal
     */
    private boolean dfs(String start, String end, List<Double> link) {
        if (start.equals(end)) { return true; }
        visited.add(start);
        HashMap<String, Double> children = nodes.get(start);
        if (children.containsKey(end)) { link.add(children.get(end)); return true; }
        for (Map.Entry<String, Double> e : children.entrySet()) {
            if (!visited.contains(e.getKey())) {
                double val = children.get(e.getKey());
                link.add(val);
                boolean isFound = dfs(e.getKey(), end, link);
                if (isFound) { return true; }
                else { link.remove(val); }
            }
        }
        return false;
    }
    
    public double[] calcEquation(String[][] equations, double[] values, String[][] queries) {
        zeros = new HashSet<>(16);
        nodes = new HashMap<>(16);
        visited = new HashSet<>(16);
        int qLen = queries.length;
        if (qLen == 0) { return new double[0]; }
        double[] ans = new double[qLen];
        addToNodes(equations, values);
        for (int i = 0; i < qLen; i++) {
            String[] query = queries[i];
            if (!nodes.containsKey(query[0]) || !nodes.containsKey(query[1]) || zeros.contains(query[1])) { 
                ans[i] = -1.0; 
            } else if (zeros.contains(query[0])) {
                ans[i] = 0;
            } else if (query[0].equals(query[1])) {
                ans[i] = 1.0;
            } else {
                if (!visited.isEmpty()) { visited.clear(); }
                List<Double> link = new ArrayList<>(); 
                boolean find = dfs(query[0], query[1], link);
                if (find) {
                    double a = 1.0;
                    for (double v : link) { a *= v; }
                    ans[i] = a;
                } else {
                    ans[i] = -1.0;
                }
            }
        }
        return ans;
    }
    
    public static void main(String args) {
        
    }
}
