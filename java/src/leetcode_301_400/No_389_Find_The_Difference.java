package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/16
 */
public class No_389_Find_The_Difference {
    public char findTheDifference(String s, String t) {
        int[] cnt = new int[26];
        int l = s.length();
        for (int i = 0; i < l; i++) {
            cnt[t.charAt(i) - 'a']++;
            cnt[s.charAt(i) - 'a']--;
        }
        cnt[t.charAt(l) - 'a']++;
        for (int i = 0; i < 26; i++) { if (cnt[i] == 1) { return (char)(i + 'a'); } }
        return '_';
    }
}
