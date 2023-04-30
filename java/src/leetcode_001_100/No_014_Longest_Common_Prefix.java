package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/5.
 * 014.最长共有前缀
 */
public class No_014_Longest_Common_Prefix {
    public String longestCommonPrefix(String[] strs) {
        int len = strs.length;
        if(len == 0) {
            return "";
        }
        if(len == 1) {
            return strs[0];
        }
        int[] lens = new int[len];
        for(int i = 0 ; i < len; i++) {
            lens[i] = strs[i].length();
        }
        StringBuilder lcp = new StringBuilder();
        int tempLen = strs[0].length();
        for(int i = 0; i < tempLen; i++) {
            char c = strs[0].charAt(i);
            boolean isCommon = true;
            for(int j = 1; j < len; j++) {
                if(i >= lens[j] || c != strs[j].charAt(i)) {
                    isCommon = false;
                    break;
                }
            }
            if(isCommon) {
                lcp.append(c);
            } else {
                break;
            }
        }
        return lcp.toString();
    }

    public static void main(String[] args) {

    }
}
