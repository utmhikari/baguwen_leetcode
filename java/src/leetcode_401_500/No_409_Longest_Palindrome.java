package leetcode_401_500;

/**
 * Created by 36191 on 2019/2/12
 */
public class No_409_Longest_Palindrome {
    public int longestPalindrome(String s) {
        int ans = 0;
        boolean[] cnt = new boolean[128];
        for (int i = 0; i < s.length(); i++) {
            int idx = s.charAt(i) - 'A';
            if (cnt[idx]) { ans += 2; cnt[idx] = false; }
            else { cnt[idx] = true; }
        }
        for (int i = 0; i < 128; i++) { if (cnt[i]) { ans++; break; } }
        return ans;
    }
}
