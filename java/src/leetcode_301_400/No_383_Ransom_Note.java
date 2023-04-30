package leetcode_301_400;

/**
 * Created by 36191 on 2019/1/16
 */
public class No_383_Ransom_Note {
    public boolean canConstruct(String ransomNote, String magazine) {
        int[] cr = new int[26], cm = new int[26];
        for (int i = 0, l = ransomNote.length(); i < l; i++) { cr[ransomNote.charAt(i) - 'a']++; }
        for (int i = 0, l = magazine.length(); i < l; i++) { cm[magazine.charAt(i) - 'a']++; }
        for (int i = 0; i < 26; i++) { if (cr[i] > cm[i]) { return false; } }
        return true;
    }
}
