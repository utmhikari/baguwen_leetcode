package leetcode_101_200;

/**
 * Created by 36191 on 2017/11/30
 * 101.对称树判断
 */

public class No_101_Symmetric_Tree {

    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }

    public boolean isSymmetric(TreeNode tn1, TreeNode tn2) {
        if(tn1 == null && tn2 == null) {
            return true;
        } else if(tn1 == null || tn2 == null) {
            return false;
        } else {
            return tn1.val == tn2.val && isSymmetric(tn1.left, tn2.right) && isSymmetric(tn1.right, tn2.left);
        }
    }

    public boolean isSymmetric(TreeNode root) {
        if(root == null) {
            return true;
        }
        return isSymmetric(root.left, root.right);
    }
}
