package leetcode_101_200;

/**
 * Created by 36191 on 2018/8/9
 */
import java.util.*;

public class No_131_Palindrome_Partitioning {
    private boolean isP(String s, int l, int r) {
        while(l < r) {
            if(s.charAt(l++) != s.charAt(r-- - 1)) { return false; }
        }
        return true;
    }

    private void visit(List<List<String>> parts, HashMap<Integer, ArrayList<Integer>> routes, List<String> part, int start,
                       String s, int l) {
        for(int right : routes.get(start)) {
            part.add(s.substring(start, right));
            if (right == l) {
                List<String> tmp = new ArrayList<String>(part);
                parts.add(tmp);
            } else {
                visit(parts, routes, part, right, s, l);
            }
            part.remove(part.size() - 1);
        }
    }

    public List<List<String>> partition(String s) {
        List<List<String>> parts = new ArrayList<>();
        int l = s.length();
        if(l == 0) { return parts; }

        HashMap<Integer, ArrayList<Integer>> routes = new HashMap<>();
        for(int i = 0; i < l; i++) {
            for(int j = i + 1; j <= l; j++) {
                if(isP(s, i, j)) {
                    if(!routes.containsKey(i)) {
                        routes.put(i, new ArrayList<Integer>());
                    }
                    routes.get(i).add(j);
                }
            }
        }
        List<String> part = new ArrayList<String>();
        visit(parts, routes, part, 0, s, l);
        return parts;
    }
}
