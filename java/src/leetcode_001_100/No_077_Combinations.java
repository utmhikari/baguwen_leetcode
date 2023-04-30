package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/21
 * 077. 组合
 */

import java.util.List;
import java.util.LinkedList;
import java.util.ArrayList;
public class No_077_Combinations {

    public void backtrack(List<List<Integer>> combinations, List<Integer> combination, int n, int k, int start) {
        if(k == 0) {
            combinations.add(new ArrayList<>(combination));
        } else {
            for(int i = start; i <= n - k + 1; i++) {
                combination.add(i);
                backtrack(combinations, combination, n, k - 1, i + 1);
                combination.remove(combination.size() - 1);
            }
        }
    }

    public List<List<Integer>> combine(int n, int k) {
        LinkedList<List<Integer>> combinations = new LinkedList<>();
        if(n == 0 || k == 0 || k > n) {
            return combinations;
        }
        ArrayList<Integer> combination = new ArrayList<>();
        backtrack(combinations, combination, n, k, 1);
        return combinations;
    }

    public void test(int n, int k) {
        List<List<Integer>> ans = combine(n, k);
        for(List<Integer> l : ans) {
            for(int i : l) {
                System.out.print(i);
                System.out.print(", ");
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        No_077_Combinations solution = new No_077_Combinations();
        solution.test(4, 2);
    }
}
