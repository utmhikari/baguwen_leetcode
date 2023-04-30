package leetcode_001_100;

/**
 * Created by 36191 on 2017/11/12
 * 096. 生成二叉搜索树数量
 */
public class No_096_Unique_Binary_Search_Trees {
    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }

    static int[] fn = new int[101];

    public int f(int n) {
        if(fn[n] != 0) {
            return fn[n];
        } else {
            int sum = 0;
            for(int i = 0; i < n; i++) {
                sum += f(i) * f(n - 1 - i);
            }
            return sum;
        }
    }

    public int numTrees(int n) {
        if(n == 0) {
            return 0;
        }
        fn[0] = 1;
        fn[1] = 1;
        fn[2] = 2;
        fn[3] = 5;
        fn[4] = 14;
        fn[5] = 42;
        fn[6] = 132;
        fn[7] = 429;
        return f(n);
    }

    public void test(int n) {
        int ans = numTrees(n);
        System.out.println(ans);
    }

    public static void main(String[] args) {
        No_096_Unique_Binary_Search_Trees solution = new No_096_Unique_Binary_Search_Trees();
        solution.test(0);
        solution.test(19);
    }
}
