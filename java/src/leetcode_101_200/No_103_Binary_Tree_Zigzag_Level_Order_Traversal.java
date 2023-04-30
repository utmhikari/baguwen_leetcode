package leetcode_101_200;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/**
 * Created by 36191 on 2017/11/30
 * 103.之字遍历
 */

public class No_103_Binary_Tree_Zigzag_Level_Order_Traversal {
    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }

    private void traverse(int level, List<List<Integer>> ans, LinkedList<TreeNode> nodes) {
        int size = nodes.size();
        if(size != 0) {
            while(size > 0) {
                TreeNode node = nodes.get(size - 1);
                if(ans.size() <= level + 1 && (node.right != null || node.left != null)) {
                    ans.add(new ArrayList<>());
                }
                if(level % 2 == 0) {
                    if(node.right != null) {
                        ans.get(level + 1).add(node.right.val);
                        nodes.add(node.right);
                    }
                    if(node.left != null) {
                        ans.get(level + 1).add(node.left.val);
                        nodes.add(node.left);
                    }
                } else {
                    if(node.left != null) {
                        ans.get(level + 1).add(node.left.val);
                        nodes.add(node.left);
                    }
                    if(node.right != null) {
                        ans.get(level + 1).add(node.right.val);
                        nodes.add(node.right);
                    }
                }
                nodes.remove(size - 1);
                size--;
            }
            traverse(level + 1, ans, nodes);
        }
    }

    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        ArrayList<List<Integer>> ans = new ArrayList<>();
        if(root == null) {
            return ans;
        }
        ArrayList<Integer> rootList = new ArrayList<Integer>(){{ add(root.val);}};
        ans.add(rootList);
        LinkedList<TreeNode> rootNode = new LinkedList<TreeNode>(){{ add(root); }};
        traverse(0, ans, rootNode);
        return ans;
    }
}
