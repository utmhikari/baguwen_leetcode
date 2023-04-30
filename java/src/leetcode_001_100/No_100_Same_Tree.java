package leetcode_001_100;

/**
 * Created by 36191 on 2017/11/30
 * 100.相同树判断
 */

public class No_100_Same_Tree {
    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }

    public boolean isSameTree(TreeNode p, TreeNode q) {
        if(p == null && q == null) return true;
        if(!(p == null || q == null)) {
            return p.val == q.val && isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
        }
        return false;
    }
}
