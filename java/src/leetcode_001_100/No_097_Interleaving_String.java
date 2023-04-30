package leetcode_001_100;

/**
 * Created by 36191 on 2017/11/14
 * 097. 插入字符串（逐个字母）
 */

public class No_097_Interleaving_String {


    public boolean isInterleave(String s1, String s2, String s3) {

        if(s1.length() + s2.length() != s3.length()) {
            return false;
        }

        boolean[][] matrix = new boolean[s2.length()+1][s1.length()+1];
        matrix[0][0] = true;
        for(int i = 1; i < matrix[0].length; i++) {
            matrix[0][i] = matrix[0][i - 1] && (s1.charAt(i - 1) == s3.charAt(i - 1));
        }

        for (int i = 1; i < matrix.length; i++) {
            matrix[i][0] = matrix[i - 1][0] && (s2.charAt(i - 1) == s3.charAt(i - 1));
        }

        for(int i = 1; i < matrix.length; i++) {
            for(int j = 1; j < matrix[0].length; j++) {
                matrix[i][j] = (matrix[i - 1][j] && (s2.charAt(i - 1) == s3.charAt(i + j - 1)))
                        || (matrix[i][j - 1] && (s1.charAt(j - 1) == s3.charAt(i + j - 1)));
            }
        }

        return matrix[s2.length()][s1.length()];
    }
/*
    public boolean isInterleave(String s1, String s2, String s3, int i1, int i2, int i3) {
        if(i3 == s3.length()) {
            return true;
        }
        boolean b1 = false;
        boolean b2 = false;
        if(i1 < s1.length()) {
            if(s1.charAt(i1) == s3.charAt(i3)) {
                b1 = isInterleave(s1, s2, s3, i1 + 1, i2, i3 + 1);
            }
        } else {
            return s2.substring(i2).equals(s3.substring(i3));
        }
        if(i2 < s2.length()) {
            if(s2.charAt(i2) == s3.charAt(i3)) {
                b2 = isInterleave(s1, s2, s3, i1, i2 + 1, i3 + 1);
            }
        } else {
            return s1.substring(i1).equals(s3.substring(i3));
        }
        return b1 || b2;
    }

    public boolean isInterleave(String s1, String s2, String s3) {
        return s1.length() + s2.length() == s3.length() && isInterleave(s1, s2, s3, 0, 0, 0);
    }
*/
    public static void main(String[] args) {
        No_097_Interleaving_String solution = new No_097_Interleaving_String();
        System.out.println(solution.isInterleave("aabcc", "dbbca", "aadbbcbcac"));
        System.out.println(solution.isInterleave("aabcc", "dbbca", "aadbbbaccc"));
    }
}
