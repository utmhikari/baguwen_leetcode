package utils;

/**
 * Created by 36191 on 2017/12/5
 * 二叉树
 */

public class TreeNode {
    public int val;
    public TreeNode left;
    public TreeNode right;
    public TreeNode(int x) { val = x; }

    /**
     * 前序遍历输出
     * @param root
     */
    public static void output(TreeNode root) {
        if (root == null) {
            System.out.print("null, ");
        } else {
            System.out.print(root.val + ", ");
            output(root.left);
            output(root.right);
        }
    }
}
