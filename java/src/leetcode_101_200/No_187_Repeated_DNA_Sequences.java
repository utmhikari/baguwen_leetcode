package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/28
 * No.187 找到所有重复过的长度为10的子串
 */

import java.util.*;

public class No_187_Repeated_DNA_Sequences {
    public List<String> findRepeatedDnaSequences(String s) {
        List<String> repl = new ArrayList<>();
        int len = s.length();
        if (len <= 10) { return repl; }
        HashSet<String> strs = new HashSet<>();
        HashSet<String> reps = new HashSet<>();
        for (int i = 0; i <= len - 10; i++) {
            String sub = s.substring(i, i + 10);
            if (strs.contains(sub)) { reps.add(sub); }
            else { strs.add(sub); }
        }
        repl.addAll(reps);
        return repl;
    }

    public static void main(String[] args) {
        No_187_Repeated_DNA_Sequences solution = new No_187_Repeated_DNA_Sequences();
        for (String s : solution.findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")) {
            System.out.println(s);
        }
    }
}
