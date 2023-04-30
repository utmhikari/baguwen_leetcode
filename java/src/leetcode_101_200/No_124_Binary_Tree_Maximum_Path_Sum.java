package leetcode_101_200;

/**
 * Created by 36191 on 2018/7/9
 */

import utils.TreeNode;

public class No_124_Binary_Tree_Maximum_Path_Sum {

    public static int m = Integer.MIN_VALUE;

    public int tmpMax(TreeNode root) {
        if (root.left == null && root.right == null) {
            m = Math.max(m, root.val);
            return root.val;
        } else if (root.left == null) {
            int tmpMax = Math.max(root.val, root.val + tmpMax(root.right));
            m = Math.max(m, Math.max(root.val, tmpMax));
            return tmpMax;
        } else if (root.right == null) {
            int tmpMax = Math.max(root.val, root.val + tmpMax(root.left));
            m = Math.max(m, Math.max(root.val, tmpMax));
            return tmpMax;
        } else {
            int l = tmpMax(root.left), r = tmpMax(root.right);
            int tmpMax = Math.max(root.val, Math.max(l, r) + root.val);
            m = Math.max(m, Math.max(Math.max(tmpMax, root.val + l + r), Math.max(l, r)));
            return tmpMax;
        }
    }

    public int maxPathSum(TreeNode root) {
        if (root == null) {
            return 0;
        }
        m = Integer.MIN_VALUE;
        tmpMax(root);
        return m;
    }

    public static void main(String[] args) {
        No_124_Binary_Tree_Maximum_Path_Sum solution = new No_124_Binary_Tree_Maximum_Path_Sum();

        TreeNode tree1 = new TreeNode(1);
        tree1.left = new TreeNode(2);
        tree1.right = new TreeNode(3);
        System.out.println(solution.maxPathSum(tree1));

        TreeNode tree2 = new TreeNode(-10);
        tree2.left = new TreeNode(9);
        tree2.right = new TreeNode(20);
        tree2.right.left = new TreeNode(15);
        tree2.right.right = new TreeNode(7);
        System.out.println(solution.maxPathSum(tree2));

        TreeNode tree3 = new TreeNode(-10);
        tree3.left = new TreeNode(-1);
        tree3.left.left = new TreeNode(-4);
        tree3.left.right = new TreeNode(-2);
        System.out.println(solution.maxPathSum(tree3));

        TreeNode tree4 = new TreeNode(5);
        tree4.left = new TreeNode(4);
        tree4.left.left = new TreeNode(11);
        tree4.left.left.left = new TreeNode(7);
        tree4.left.left.right = new TreeNode(2);
        tree4.right = new TreeNode(8);
        tree4.right.left = new TreeNode(13);
        tree4.right.right = new TreeNode(4);
        tree4.right.right.right = new TreeNode(1);
        System.out.println(solution.maxPathSum(tree4));

    }
}
