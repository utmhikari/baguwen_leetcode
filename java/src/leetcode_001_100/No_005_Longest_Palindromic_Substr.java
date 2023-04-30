package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/3.
 * 005.最长回文子串——Manacher算法
 */

public class No_005_Longest_Palindromic_Substr {

    /**
     * 编码成  $#a#a#b#b#\0  的样式
     * 就可以考虑以单个字符或是两个相同字符为中心的情况
     * @param s
     * @return
     */
    public static String pre(String s) {
        StringBuilder preStr = new StringBuilder();
        int length = s.length();
        preStr.append('$');
        for(int i = 0; i < length; i++)
        {
            preStr.append('#');
            preStr.append(s.charAt(i));
        }
        preStr.append('#');
        return preStr.toString();
    }

    public String longestPalindrome(String s) {
        String preStr = pre(s);
        int mx = 0;  // 右边界
        int pi = 1;  // 对称中心
        int len = preStr.length();
        int[] p = new int[len];
        for(int i = 1; i < len; i++)
        {
            // 如果右边界大于i则为第1个开始
            if(mx > i) {
                int size1 = mx - i;
                int size2 = p[2 * pi - i];
                if(size1 < size2) {
                    p[i] = size1;
                } else {
                    p[i] = size2;
                }
            } else {
                p[i] = 1;
            }
            while(i - p[i] > 0 && i + p[i] < len && preStr.charAt(i - p[i]) == preStr.charAt(i + p[i])) {
                p[i]++;
            }
            if(i + p[i] > mx) {
                mx = p[i] + i;
                pi = i;
            }
        }
        int maxlen = 0; // 最大边界长度（mx + 1）
        int index = 0; // 最长回文串中心
        for(int i = 0; i < len; i++)
        {
            if(p[i] > maxlen)
            {
                maxlen = p[i];
                index = i;
            }
        }
        int startIndex = index - maxlen + 1;
        int endIndex = index + maxlen - 1;
        StringBuilder lp = new StringBuilder();
        for(int i = startIndex; i <= endIndex; i++) {
            char c = preStr.charAt(i);
            if(c != '#') {
                lp.append(c);
            }
        }
        return lp.toString();

    }

    public static void main(String[] args) {
        No_005_Longest_Palindromic_Substr solution = new No_005_Longest_Palindromic_Substr();
        String a = "babad";
        String b = "cbbe";
        System.out.println(solution.longestPalindrome(a));
        System.out.println(solution.longestPalindrome(b));
    }
}
