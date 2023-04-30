package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/16
 */
public class No_387_First_Unique_Character_in_a_String {
    public int firstUniqChar(String s) {
        int l = s.length();
        if (l <= 1) { return l == 1 ? 0 : -1; }
        boolean[] ndups = new boolean[l];
        int[] marks = new int[26];
        for (int i = 0; i < 26; i++) { marks[i] = -1; }
        for (int i = 0; i < l; i++) {
            int idx = s.charAt(i) - 'a';
            if (marks[idx] == -1) { marks[idx] = i; ndups[i] = true;}
            else if (marks[idx] >= 0) { ndups[marks[idx]] = false; marks[idx] = -2; ndups[i] = false; }
        }
        for (int i = 0; i < l; i++) { if (ndups[i]) { return i; } }
        return -1;
    }
}
