package leetcode_301_400;

/**
 * Created by 36191 on 2019/2/12
 */
public class No_395_Longest_Substring_with_At_Least_K_Repeating_Characters {
    public int longestSubstring(String s, int k) {
        return ls(s, 0, s.length() - 1, k);
    }

    private int ls(String s, int start, int end, int k) {
        if (end - start + 1 < k) { return 0; }
        int[] cnt = new int[26];
        for (int i = start; i <= end; i++) { cnt[s.charAt(i) - 'a']++; }
        for (int i = 0; i < 26; i++) {
            if (cnt[i] > 0 && cnt[i] < k) {
                for (int j = start; j <= end; j++) {
                    if (s.charAt(j) == i + 'a') {
                        return Math.max(ls(s, start, j - 1, k), ls(s, j + 1, end, k));
                    }
                }
            }
        }
        return end - start + 1;
    }
}
