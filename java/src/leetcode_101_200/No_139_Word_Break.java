package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/19
 */
import java.util.*;

public class No_139_Word_Break {

    public boolean wordBreak(String s, List<String> wordDict) {
        int len = s.length();
        boolean[] bools = new boolean[len + 1];
        bools[0] = true;
        for(int i = 0; i < s.length(); i++) {
            if (bools[i]) {
                for (String word : wordDict) {
                    int right = word.length() + i;
                    if (right <= len && !bools[right] && word.equals(s.substring(i, right))) {
                        bools[right] = true;
                    }
                }
            }
        }
        return bools[len];
    }

    public static void main(String[] args) {
        No_139_Word_Break solution = new No_139_Word_Break();
        String s1 = "leetcode", s2 = "applepenapple", s3 = "catsandog";
        List<String> wd1 = Arrays.asList("leet", "code");
        List<String> wd2 = Arrays.asList("apple", "pen");
        List<String> wd3 = Arrays.asList("cats", "dog", "sand", "and", "cat");
        System.out.println(solution.wordBreak(s1, wd1));
        System.out.println(solution.wordBreak(s2, wd2));
        System.out.println(solution.wordBreak(s3, wd3));
    }
}
