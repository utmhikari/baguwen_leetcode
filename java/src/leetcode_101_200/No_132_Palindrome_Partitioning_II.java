package leetcode_101_200;

/**
 * Created by 36191 on 2019/4/27
 */
public class No_132_Palindrome_Partitioning_II {

    private void output(int[][] g) {
        int l = g.length;
        System.out.println("------------------------------------\n");
        for (int i = 0; i < l; ++i) {
            for (int j = 0; j < l; ++j) {
                System.out.print(g[i][j] + ", ");
            }
            System.out.println();
        }
        System.out.println("\n------------------------------------");
    }

    public int minCut(String s) {
        int l = s.length();
        if (l <= 1) { return 0; }
        int[][] g = new int[l][l];
        char[] cs = s.toCharArray();
        // tag 1 char palindromes
        for (int i = 0; i < l; ++i) {
            for (int j = i; j < l; ++j) {
                g[i][j] = j - i + 1;
            }
        }
        // based on 1 char
        int r;
        for (int i = 0; i < l; ++i) {
            r = 1;
            while (i - r >= 0 && i + r < l && cs[i - r] == cs[i + r]) {
                g[i - r][i + r] = 1;
                ++r;
            }
        }
        // based on 2 chars
        for (int i = 0; i < l - 1; ++i) {
            if (cs[i] == cs[i + 1]) {
                g[i][i + 1] = 1;
                r = 1;
                while (i - r >= 0 && i + r + 1 < l && cs[i - r] == cs[i + r + 1]) {
                    g[i - r][i + r + 1] = 1;
                    ++r;
                }
            }
        }
        // traverse graph
        for (int j = 1; j < l; ++j) {
            for (int i = j - 1; i >= 0; --i) {
                if (g[i][j] != 1) {
                    for (int k = 1; k <= j - i; ++k) {
                        if (g[i][j - k] < j - i - k + 1) {
                            g[i][j] = Math.min(g[i][j], g[i][j - k] + g[j - k + 1][j]);
                        }
                        if (g[i + k][j] < j - i - k + 1) {
                            g[i][j] = Math.min(g[i][j], g[i][i + k - 1] + g[i + k][j]);
                        }
                    }
                }
            }
        }
        return g[0][l - 1] - 1;
    }

    public static void main(String[] args) {
        No_132_Palindrome_Partitioning_II solution = new No_132_Palindrome_Partitioning_II();
        String[] inputs = new String[] {
                "abccccb",
        };
        for (String s : inputs) {
            System.out.println(solution.minCut(s));
        }
    }
}
