package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/24
 */


public class No_168_Excel_Sheet_Column_Title {
    public String convertToTitle(int n) {
        char c = 'A';
        StringBuilder t = new StringBuilder();
        while (n > 26) {
            t.append((char)(c + ((n - 1) % 26)));
            n = (n - 1) / 26;
        }
        t.append((char)(c + ((n - 1) % 26)));
        return t.reverse().toString();
    }

    public static void main(String[] args) {
        No_168_Excel_Sheet_Column_Title solution = new No_168_Excel_Sheet_Column_Title();
        int[] inputs = new int[] {1, 28, 701};
        for (int i : inputs) {
            System.out.println(solution.convertToTitle(i));
        }
    }
}
