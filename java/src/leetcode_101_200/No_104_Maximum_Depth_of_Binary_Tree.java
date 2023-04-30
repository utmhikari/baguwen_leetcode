package leetcode_101_200;

/**
 * Created by 36191 on 2017/12/1
 * 104.二叉树深度
 */

public class No_104_Maximum_Depth_of_Binary_Tree {
    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }

    private static int maxDepth;

    private void dfs(TreeNode node, int depth) {
        if(node != null) {
            if(depth > maxDepth) {
                maxDepth = depth;
            }
            dfs(node.left, depth + 1);
            dfs(node.right, depth + 1);
        }
    }

    public int maxDepth(TreeNode root) {
        maxDepth = 0;
        dfs(root, 1);
        return maxDepth;
    }
}
