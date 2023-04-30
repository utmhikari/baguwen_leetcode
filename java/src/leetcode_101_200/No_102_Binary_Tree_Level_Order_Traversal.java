package leetcode_101_200;

import java.util.List;
import java.util.ArrayList;

/**
 * Created by 36191 on 2017/11/30
 * 102.按层遍历二叉树
 */

public class No_102_Binary_Tree_Level_Order_Traversal {

    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }

    private void traverse(TreeNode node, int level, List<List<Integer>> ans) {
        if(node != null) {
            if(ans.size() <= level) {
                ans.add(new ArrayList<>());
            }
            ans.get(level).add(node.val);
            traverse(node.left, level + 1, ans);
            traverse(node.right, level + 1, ans);
        }
    }

    public List<List<Integer>> levelOrder(TreeNode root) {
        ArrayList<List<Integer>> ans = new ArrayList<>();
        traverse(root, 0, ans);
        return ans;
    }
}
